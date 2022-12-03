package accountpg

import (
	"errors"
	"gorm.io/gorm"
)

const (
	initialAccountType    = "Main"
	initialAccountName    = "General"
	initialAccountBalance = 10000

	SecondaryAccountType  = "Secondary"
	DefaultAccountBalance = 0
)

type PostgresAccountStorage struct {
	db *gorm.DB
}

func NewPostgresAccountStorage(db *gorm.DB) *PostgresAccountStorage {
	return &PostgresAccountStorage{db: db}
}

func (s *PostgresAccountStorage) CreateInitialAccount(userId uint) error {
	return s.createInitialAccountImpl(userId)
}

func (s *PostgresAccountStorage) CreateCustomAccount(userId uint, accountName string) error {
	return s.createCustomAccountImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccount(userId uint, accountName string) (*Account, error) {
	return s.getAccountImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccounts(userId uint) ([]*Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *PostgresAccountStorage) UpdateBalance(userId uint, accountName string, balance int) error {
	return s.updateBalanceImpl(userId, accountName, balance)
}

func (s *PostgresAccountStorage) DeleteAccount(userId uint, accountName string) error {
	return s.deleteAccountImpl(userId, accountName)
}

func (s *PostgresAccountStorage) createInitialAccountImpl(userId uint) error {
	var account Account

	account.UserId = userId
	account.Type = initialAccountType
	account.Name = initialAccountName
	account.Balance = initialAccountBalance

	err := s.db.Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresAccountStorage) createCustomAccountImpl(userId uint, accountName string) error {
	var account Account

	account.UserId = userId
	account.Type = SecondaryAccountType
	account.Name = accountName
	account.Balance = DefaultAccountBalance

	err := s.db.Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresAccountStorage) updateBalanceImpl(userId uint, accountName string, balance int) error {
	account, err := s.getAccountImpl(userId, accountName)
	if err != nil {
		return err
	}

	account.Balance = balance
	err = s.db.Save(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresAccountStorage) deleteAccountImpl(userId uint, accountName string) error {
	account, err := s.getAccountImpl(userId, accountName)
	if err != nil {
		return err
	}

	if s.isMainAccountType(account.Type) {
		return errors.New("cannot delete account with main type")
	}

	if !s.isZeroAccountBalance(account.Balance) {
		return errors.New("cannot delete account with non-zero balance")
	}

	err = s.db.Delete(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresAccountStorage) getAccountImpl(userId uint, accountName string) (*Account, error) {
	var account Account

	err := s.db.Table("accounts").Where("user_id = ? AND name = ?", userId, accountName).First(&account).Error
	if err != nil {
		// log.Error(err.Error())
		return nil, errors.New("cannot find account")
	}

	return &account, nil
}

func (s *PostgresAccountStorage) getAccountsImpl(userId uint) ([]*Account, error) {
	var accounts []*Account

	err := s.db.Table("accounts").Where("user_id = ?", userId).Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *PostgresAccountStorage) isZeroAccountBalance(balance int) bool {
	return balance <= 0
}

func (s *PostgresAccountStorage) isMainAccountType(accountType string) bool {
	return accountType == "Main"
}
