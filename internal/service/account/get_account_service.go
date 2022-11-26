package account

import (
	domain "github.com/vaberof/banking_app/internal/domain/account"
)

type GetAccountService struct {
	accountStorage AccountStorage
}

func NewGetAccountService(accountStorage AccountStorage) *GetAccountService {
	return &GetAccountService{accountStorage: accountStorage}
}

func (s *GetAccountService) CreateInitialAccount(userId uint) error {
	return s.accountStorage.CreateInitialAccount(userId)
}

func (s *GetAccountService) CreateCustomAccount(userId uint, accountName string) error {
	return s.accountStorage.CreateCustomAccount(userId, accountName)
}

func (s *GetAccountService) UpdateBalance(userId uint, accountName string, balance int) error {
	return s.accountStorage.UpdateBalance(userId, accountName, balance)
}

func (s *GetAccountService) DeleteAccount(userId uint, accountName string) error {
	return s.accountStorage.DeleteAccount(userId, accountName)
}

func (s *GetAccountService) GetAccount(userId uint, accountName string) (*GetAccount, error) {
	return s.getAccountImpl(userId, accountName)
}

func (s *GetAccountService) GetAccounts(userId uint) ([]*GetAccount, error) {
	return s.getAccountsImpl(userId)
}

func (s *GetAccountService) getAccountImpl(userId uint, accountName string) (*GetAccount, error) {
	domainAccount, err := s.accountStorage.GetAccount(userId, accountName)
	if err != nil {
		return nil, err
	}

	getAccount := s.domainAccountToGetAccount(domainAccount)
	return getAccount, nil
}

func (s *GetAccountService) getAccountsImpl(userId uint) ([]*GetAccount, error) {
	infraAccounts, err := s.accountStorage.GetAccounts(userId)
	if err != nil {
		return nil, err
	}

	domainAccounts := s.domainAccountsToGetAccounts(infraAccounts)
	return domainAccounts, nil
}

func (s *GetAccountService) domainAccountToGetAccount(domainAccount *domain.Account) *GetAccount {
	var account GetAccount

	account.UserId = domainAccount.UserId
	account.Type = domainAccount.Type
	account.Name = domainAccount.Name
	account.Balance = domainAccount.Balance

	return &account
}

func (s *GetAccountService) domainAccountsToGetAccounts(domainAccounts []*domain.Account) []*GetAccount {
	var getAccounts []*GetAccount

	for i := 0; i < len(domainAccounts); i++ {
		domainAccount := domainAccounts[i]
		getAccounts = append(getAccounts, s.domainAccountToGetAccount(domainAccount))
	}

	return getAccounts
}
