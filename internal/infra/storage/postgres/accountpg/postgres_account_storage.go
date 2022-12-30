package accountpg

import (
	"errors"
	"github.com/sirupsen/logrus"
	domain "github.com/vaberof/MockBankingApplication/internal/domain/account"
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

func (s *PostgresAccountStorage) CreateCustomAccount(userId uint, accountName string) (*domain.Account, error) {
	return s.createCustomAccountImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccountByName(userId uint, accountName string) (*domain.Account, error) {
	return s.getAccountByNameImpl(userId, accountName)
}

func (s *PostgresAccountStorage) GetAccountById(accountId uint) (*domain.Account, error) {
	return s.getAccountByIdImpl(accountId)
}

func (s *PostgresAccountStorage) GetAccounts(userId uint) ([]*domain.Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *PostgresAccountStorage) UpdateBalance(postgresAccount *PostgresAccount, balance int) error {
	return s.updateBalanceImpl(postgresAccount, balance)
}

func (s *PostgresAccountStorage) DeleteAccount(domainAccount *domain.Account) error {
	return s.deleteAccountImpl(domainAccount)
}

func (s *PostgresAccountStorage) createInitialAccountImpl(userId uint) error {
	var postgresAccount PostgresAccount

	postgresAccount.UserId = userId
	postgresAccount.Type = initialAccountType
	postgresAccount.Name = initialAccountName
	postgresAccount.Balance = initialAccountBalance

	err := s.db.Table("accounts").Create(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "createInitialAccountImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresAccountStorage) createCustomAccountImpl(userId uint, accountName string) (*domain.Account, error) {
	var postgresAccount PostgresAccount

	postgresAccount.UserId = userId
	postgresAccount.Type = SecondaryAccountType
	postgresAccount.Name = accountName
	postgresAccount.Balance = SecondaryAccountBalance

	err := s.db.Table("accounts").Create(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "createCustomAccountImpl",
		}).Error(err)

		return nil, err
	}

	domainAccount := BuildDomainAccount(&postgresAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountByNameImpl(userId uint, accountName string) (*domain.Account, error) {
	var postgresAccount PostgresAccount

	err := s.db.Table("accounts").Where("user_id = ? AND name = ?", userId, accountName).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "getAccountByNameImpl",
		}).Error(err)

		return nil, errors.New("cannot find account")
	}

	domainAccount := BuildDomainAccount(&postgresAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountByIdImpl(accountId uint) (*domain.Account, error) {
	var postgresAccount PostgresAccount

	err := s.db.Table("accounts").Where("id = ?", accountId).First(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "getAccountByIdImpl",
		}).Error(err)

		return nil, errors.New("cannot find account")
	}

	domainAccount := BuildDomainAccount(&postgresAccount)

	return domainAccount, nil
}

func (s *PostgresAccountStorage) getAccountsImpl(userId uint) ([]*domain.Account, error) {
	var postgresAccounts []*PostgresAccount

	err := s.db.Table("accounts").Where("user_id = ?", userId).Find(&postgresAccounts).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "getAccountsImpl",
		}).Error(err)

		return nil, err
	}

	domainAccounts := BuildDomainAccounts(postgresAccounts)
	return domainAccounts, nil
}

func (s *PostgresAccountStorage) updateBalanceImpl(postgresAccount *PostgresAccount, balance int) error {
	postgresAccount.Balance = balance

	err := s.db.Table("accounts").Save(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "updateBalanceImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresAccountStorage) deleteAccountImpl(domainAccount *domain.Account) error {
	postgresAccount := BuildPostgresAccount(domainAccount)

	err := s.db.Table("accounts").Delete(&postgresAccount).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer":   "infra",
			"package": "accountpg",
			"func":    "deleteAccountImpl",
		}).Error(err)

		return err
	}

	return nil
}
