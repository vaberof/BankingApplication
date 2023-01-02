package transferpg

import (
	"github.com/sirupsen/logrus"
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/accountpg"
	service "github.com/vaberof/MockBankingApplication/internal/service/transfer"
	"gorm.io/gorm"
)

type PostgresTransferStorage struct {
	accountStorage AccountStorage
	db             *gorm.DB
}

func NewPostgresTransferStorage(db *gorm.DB, accountStorage AccountStorage) *PostgresTransferStorage {
	return &PostgresTransferStorage{
		db:             db,
		accountStorage: accountStorage,
	}
}

func (s *PostgresTransferStorage) SaveTransfer(
	senderUsername string,
	senderAccount *account.Account,
	payeeUsername string,
	payeeAccount *account.Account,
	amount uint,
	transferType string) (*service.Transfer, error) {

	return s.saveTransferImpl(senderUsername, senderAccount, payeeUsername, payeeAccount, amount, transferType)
}

func (s *PostgresTransferStorage) GetTransfers(userId uint) ([]*service.Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *PostgresTransferStorage) saveTransferImpl(
	senderUsername string,
	senderAccount *account.Account,
	payeeUsername string,
	payeeAccount *account.Account,
	amount uint,
	transferType string) (*service.Transfer, error) {

	postgresSenderAccount, postgresPayeeAccount, err := s.preprocessTransfer(senderAccount, payeeAccount)
	if err != nil {
		return nil, err
	}

	err = s.processTransfer(postgresSenderAccount, postgresPayeeAccount, amount)
	if err != nil {
		return nil, err
	}

	serviceTransfer, err := s.createTransfer(
		postgresSenderAccount.UserId,
		senderUsername,
		postgresSenderAccount.Id,
		postgresPayeeAccount.UserId,
		payeeUsername,
		payeeAccount.Id,
		amount,
		transferType)
	if err != nil {
		return nil, err
	}

	return serviceTransfer, nil
}

func (s *PostgresTransferStorage) preprocessTransfer(
	senderAccount *account.Account,
	payeeAccount *account.Account) (*accountpg.PostgresAccount, *accountpg.PostgresAccount, error) {

	senderPostgresAccount := accountpg.BuildPostgresAccount(senderAccount)
	payeePostgresAccount := accountpg.BuildPostgresAccount(payeeAccount)

	return senderPostgresAccount, payeePostgresAccount, nil
}

func (s *PostgresTransferStorage) processTransfer(senderAccount *accountpg.PostgresAccount, payeeAccount *accountpg.PostgresAccount, amount uint) error {
	convAmount := int(amount)

	err := s.accountStorage.UpdateBalance(senderAccount, senderAccount.Balance-convAmount)
	if err != nil {
		return err
	}

	err = s.accountStorage.UpdateBalance(payeeAccount, payeeAccount.Balance+convAmount)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresTransferStorage) createTransfer(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	transferType string) (*service.Transfer, error) {

	var postgresTransfer PostgresTransfer

	postgresTransfer.SenderId = senderId
	postgresTransfer.SenderUsername = senderUsername
	postgresTransfer.SenderAccountId = senderAccountId
	postgresTransfer.PayeeId = payeeId
	postgresTransfer.PayeeUsername = payeeUsername
	postgresTransfer.PayeeAccountId = payeeAccountId
	postgresTransfer.Amount = amount
	postgresTransfer.TransferType = transferType

	err := s.db.Table("transfers").Create(&postgresTransfer).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "transferpg",
			"func":    "createTransfer",
		}).Error(err)

		return nil, err
	}

	serviceTransfer := BuildServiceTransfer(&postgresTransfer)
	return serviceTransfer, nil
}

func (s *PostgresTransferStorage) getTransfersImpl(userId uint) ([]*service.Transfer, error) {
	var postgresTransfers []*PostgresTransfer

	err := s.db.Table("transfers").Where("sender_id = ?", userId).Find(&postgresTransfers).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "transferpg",
			"func":    "getTransfersImpl",
		}).Error(err)

		return nil, err
	}

	serviceTransfers := BuildServiceTransfers(postgresTransfers)

	return serviceTransfers, nil
}
