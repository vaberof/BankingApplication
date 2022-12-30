package accountpg

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
	"gorm.io/gorm"
)

const (
	initialAccountType      = "Main"
	SecondaryAccountType    = "Secondary"
	initialAccountName      = "General"
	initialAccountBalance   = 10000
	SecondaryAccountBalance = 0
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

func (s *PostgresAccountStorage) CreateCustomAccount(userId uint, accountName string) (*account.Account, error) {
	return s.createCustomAccountImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccountByName(userId uint, accountName string) (*account.Account, error) {
	return s.getAccountByNameImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccountById(accountId uint) (*account.Account, error) {
	return s.getAccountByIdImpl(accountId)
}

func (s *PostgresAccountStorage) GetAccounts(userId uint) ([]*account.Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *PostgresAccountStorage) UpdateBalance(infraAccount *Account, balance int) error {
	return s.updateBalanceImpl(infraAccount, balance)
}

func (s *PostgresAccountStorage) DeleteAccount(domainAccount *account.Account) error {
	return s.deleteAccountImpl(domainAccount)
}

func (s *PostgresAccountStorage) createInitialAccountImpl(userId uint) error {
	var account Account

	account.UserId = userId
	account.Type = initialAccountType
	account.Name = initialAccountName
	account.Balance = initialAccountBalance

	err := s.db.Create(&account).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "createInitialAccountImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresAccountStorage) createCustomAccountImpl(userId uint, accountName string) (*account.Account, error) {
	var infraAccount Account

	infraAccount.UserId = userId
	infraAccount.Type = SecondaryAccountType
	infraAccount.Name = accountName
	infraAccount.Balance = SecondaryAccountBalance

	err := s.db.Create(&infraAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "createCustomAccountImpl",
		}).Error(err)

		return nil, err
	}

	domainAccount := s.infraAccountToDomain(&infraAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountByNameImpl(userId uint, accountName string) (*account.Account, error) {
	var infraAccount Account

	err := s.db.Table("accounts").Where("user_id = ? AND name = ?", userId, accountName).First(&infraAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "getAccountByNameImpl",
		}).Error(err)

		return nil, errors.New("cannot find account")
	}

	domainAccount := s.infraAccountToDomain(&infraAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountByIdImpl(accountId uint) (*account.Account, error) {
	var infraAccount Account

	err := s.db.Table("accounts").Where("id = ?", accountId).First(&infraAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "getAccountByIdImpl",
		}).Error(err)

		return nil, errors.New("cannot find account")
	}

	domainAccount := s.infraAccountToDomain(&infraAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountsImpl(userId uint) ([]*account.Account, error) {
	var infraAccounts []*Account

	err := s.db.Table("accounts").Where("user_id = ?", userId).Find(&infraAccounts).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "getAccountsImpl",
		}).Error(err)

		return nil, err
	}

	domainAccounts := s.infraAccountsToDomain(infraAccounts)
	return domainAccounts, nil
}

func (s *PostgresAccountStorage) updateBalanceImpl(account *Account, balance int) error {
	account.Balance = balance

	err := s.db.Save(&account).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "updateBalanceImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresAccountStorage) deleteAccountImpl(domainAccount *account.Account) error {
	infraAccount := s.DomainAccountToInfra(domainAccount)

	err := s.db.Delete(&infraAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "deleteAccountImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresAccountStorage) isZeroAccountBalance(balance int) bool {
	return balance <= 0
}

func (s *PostgresAccountStorage) isMainAccountType(accountType string) bool {
	return accountType == "Main"
}
