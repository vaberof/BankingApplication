package transferpg

import (
	"errors"
	"github.com/sirupsen/logrus"
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
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	transferType string) (*service.Transfer, error) {

	return s.saveTransferImpl(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount, transferType)
}

func (s *PostgresTransferStorage) GetTransfers(userId uint) ([]*service.Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *PostgresTransferStorage) saveTransferImpl(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	transferType string) (*service.Transfer, error) {

	senderAccount, payeeAccount, err := s.preprocessTransfer(senderAccountId, payeeAccountId, amount)
	if err != nil {
		return nil, err
	}

	err = s.processTransfer(senderAccount, payeeAccount, amount)
	if err != nil {
		return nil, err
	}

	serviceTransfer, err := s.createTransfer(
		senderId,
		senderUsername,
		senderAccountId,
		payeeId,
		payeeUsername,
		payeeAccountId,
		amount,
		transferType)
	if err != nil {
		return nil, err
	}

	return serviceTransfer, nil
}

func (s *PostgresTransferStorage) preprocessTransfer(
	senderAccountId uint,
	payeeAccountId uint,
	amount uint) (*accountpg.PostgresAccount, *accountpg.PostgresAccount, error) {

	senderAccount, err := s.accountStorage.GetAccountById(senderAccountId)
	if err != nil {
		return nil, nil, err
	}

	if !s.isEnoughFunds(senderAccount.Balance, amount) {
		return nil, nil, errors.New("insufficient funds")
	}

	payeeAccount, err := s.accountStorage.GetAccountById(payeeAccountId)
	if err != nil {
		return nil, nil, err
	}

	senderInfraAccount := accountpg.BuildPostgresAccount(senderAccount)
	payeeAccountInfraAccount := accountpg.BuildPostgresAccount(payeeAccount)

	return senderInfraAccount, payeeAccountInfraAccount, nil
}

func (s *PostgresTransferStorage) processTransfer(senderAccount *accountpg.PostgresAccount, payeeAccount *accountpg.PostgresAccount, amount uint) error {
	err := s.accountStorage.UpdateBalance(senderAccount, senderAccount.Balance-int(amount))
	if err != nil {
		return err
	}

	err = s.accountStorage.UpdateBalance(payeeAccount, payeeAccount.Balance+int(amount))
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
			"layer": "infra",
			"func":  "createTransfer",
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
			"layer": "infra",
			"func":  "getTransfersImpl",
		}).Error(err)

		return nil, err
	}

	serviceTransfers := BuildServiceTransfers(postgresTransfers)

	return serviceTransfers, nil
}

func (s *PostgresTransferStorage) isEnoughFunds(balance int, amount uint) bool {
	return balance-int(amount) >= 0
}
