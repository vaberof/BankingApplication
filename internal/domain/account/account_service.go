package account

import (
	infra "github.com/vaberof/banking_app/internal/infra/storage/postgres/accountpg"
)

type AccountService struct {
	accountStorage AccountStorage
}

func NewAccountService(accountStorage AccountStorage) *AccountService {
	return &AccountService{accountStorage: accountStorage}
}

func (s *AccountService) CreateInitialAccount(userId uint) error {
	return s.accountStorage.CreateInitialAccount(userId)
}

func (s *AccountService) CreateCustomAccount(userId uint, accountName string) error {
	return s.accountStorage.CreateCustomAccount(userId, accountName)
}

func (s *AccountService) UpdateBalance(userId uint, accountName string, balance int) error {
	return s.accountStorage.UpdateBalance(userId, accountName, balance)
}

func (s *AccountService) DeleteAccount(userId uint, accountName string) error {
	return s.accountStorage.DeleteAccount(userId, accountName)
}

func (s *AccountService) GetAccount(userId uint, accountName string) (*Account, error) {
	return s.getAccountImpl(userId, accountName)
}

func (s *AccountService) GetAccounts(userId uint) ([]*Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *AccountService) getAccountImpl(userId uint, accountName string) (*Account, error) {
	infraAccount, err := s.accountStorage.GetAccount(userId, accountName)
	if err != nil {
		return nil, err
	}

	domainAccount := s.infraAccountToDomain(infraAccount)

	return domainAccount, nil
}

func (s *AccountService) getAccountsImpl(userId uint) ([]*Account, error) {
	infraAccounts, err := s.accountStorage.GetAccounts(userId)
	if err != nil {
		return nil, err
	}

	domainAccounts := s.infraAccountsToDomain(infraAccounts)

	return domainAccounts, nil
}

func (s *AccountService) infraAccountToDomain(infraAccount *infra.Account) *Account {
	var account Account

	account.UserId = infraAccount.UserId
	account.Type = infraAccount.Type
	account.Name = infraAccount.Name
	account.Balance = infraAccount.Balance

	return &account
}

func (s *AccountService) infraAccountsToDomain(infraAccounts []*infra.Account) []*Account {
	var accounts []*Account

	for i := 0; i < len(infraAccounts); i++ {
		infraAccount := infraAccounts[i]
		accounts = append(accounts, s.infraAccountToDomain(infraAccount))
	}

	return accounts
}
