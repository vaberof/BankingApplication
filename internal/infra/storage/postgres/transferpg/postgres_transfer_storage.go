package transferpg

import (
	"errors"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/accountpg"
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

func (s *PostgresTransferStorage) MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint, transferType string) error {
	return s.makeTransferImpl(senderId, senderAccountId, payeeId, payeeAccountId, amount, transferType)
}

func (s *PostgresTransferStorage) GetTransfers(userId uint) []*Transfer {
	return s.getTransfersImpl(userId)
}

func (s *PostgresTransferStorage) makeTransferImpl(
	senderId uint,
	senderAccountId uint,
	payeeId uint,
	payeeAccountId uint,
	amount uint,
	transferType string) error {

	senderAccount, payeeAccount, err := s.preprocessTransfer(senderId, senderAccountId, payeeId, payeeAccountId, amount)
	if err != nil {
		return err
	}

	err = s.processTransfer(senderAccount, payeeAccount, amount)
	if err != nil {
		return err
	}

	err = s.saveTransfer(senderId, senderAccountId, payeeId, payeeAccountId, amount, transferType)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresTransferStorage) preprocessTransfer(
	senderId uint,
	senderAccountId uint,
	payeeId uint,
	payeeAccountId uint,
	amount uint) (*accountpg.Account, *accountpg.Account, error) {

	senderAccount, err := s.accountStorage.GetAccountById(senderId, senderAccountId)
	if err != nil {
		return nil, nil, err
	}

	if !s.isEnoughFunds(senderAccount.Balance, amount) {
		return nil, nil, errors.New("insufficient funds")
	}

	payeeAccount, err := s.accountStorage.GetAccountById(payeeId, payeeAccountId)
	if err != nil {
		return nil, nil, err
	}

	senderInfraAccount := s.accountStorage.DomainAccountToInfra(senderAccount)
	payeeAccountInfraAccount := s.accountStorage.DomainAccountToInfra(payeeAccount)

	return senderInfraAccount, payeeAccountInfraAccount, nil
}

func (s *PostgresTransferStorage) processTransfer(senderAccount *accountpg.Account, payeeAccount *accountpg.Account, amount uint) error {
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

func (s *PostgresTransferStorage) isEnoughFunds(balance int, amount uint) bool {
	return balance-int(amount) >= 0
}

func (s *PostgresTransferStorage) saveTransfer(
	senderId uint,
	senderAccountId uint,
	payeeId uint,
	payeeAccountId uint,
	amount uint,
	transferType string) error {

	var transfer Transfer

	transfer.SenderId = senderId
	transfer.SenderAccountId = senderAccountId
	transfer.PayeeId = payeeId
	transfer.PayeeAccountId = payeeAccountId
	transfer.Amount = amount
	transfer.TransferType = transferType

	err := s.db.Create(&transfer).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresTransferStorage) getTransfersImpl(userId uint) []*Transfer {
	return nil
}
