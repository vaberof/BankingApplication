package account

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/account"
)

type GetAccountResponseService struct {
	accountStorage AccountStorage
}

func NewGetAccountResponseService(accountStorage AccountStorage) *GetAccountResponseService {
	return &GetAccountResponseService{accountStorage: accountStorage}
}

func (s *GetAccountResponseService) CreateInitialAccount(userId uint) error {
	return s.accountStorage.CreateInitialAccount(userId)
}

func (s *GetAccountResponseService) CreateCustomAccount(userId uint, accountName string) error {
	return s.accountStorage.CreateCustomAccount(userId, accountName)
}

func (s *GetAccountResponseService) DeleteAccount(userId uint, accountName string) error {
	return s.accountStorage.DeleteAccount(userId, accountName)
}

func (s *GetAccountResponseService) GetAccountByName(userId uint, accountName string) (*GetAccountResponse, error) {
	return s.getAccountByNameImpl(userId, accountName)
}

func (s *GetAccountResponseService) GetAccountById(accountId uint) (*GetAccountResponse, error) {
	return s.getAccountByIdImpl(accountId)
}

func (s *GetAccountResponseService) GetAccounts(userId uint) ([]*GetAccountResponse, error) {
	return s.getAccountsImpl(userId)
}

func (s *GetAccountResponseService) getAccountByNameImpl(userId uint, accountName string) (*GetAccountResponse, error) {
	domainAccount, err := s.accountStorage.GetAccountByName(userId, accountName)
	if err != nil {
		return nil, err
	}

	getAccount := s.domainAccountToGetAccount(domainAccount)
	return getAccount, nil
}

func (s *GetAccountResponseService) getAccountByIdImpl(accountId uint) (*GetAccountResponse, error) {
	domainAccount, err := s.accountStorage.GetAccountById(accountId)
	if err != nil {
		return nil, err
	}

	getAccount := s.domainAccountToGetAccount(domainAccount)
	return getAccount, nil
}

func (s *GetAccountResponseService) getAccountsImpl(userId uint) ([]*GetAccountResponse, error) {
	infraAccounts, err := s.accountStorage.GetAccounts(userId)
	if err != nil {
		return nil, err
	}

	domainAccounts := s.domainAccountsToGetAccounts(infraAccounts)
	return domainAccounts, nil
}

func (s *GetAccountResponseService) domainAccountToGetAccount(domainAccount *domain.Account) *GetAccountResponse {
	var account GetAccountResponse

	account.Id = domainAccount.Id
	account.Type = domainAccount.Type
	account.Name = domainAccount.Name
	account.Balance = domainAccount.Balance

	return &account
}

func (s *GetAccountResponseService) domainAccountsToGetAccounts(domainAccounts []*domain.Account) []*GetAccountResponse {
	var getAccounts []*GetAccountResponse

	for i := 0; i < len(domainAccounts); i++ {
		domainAccount := domainAccounts[i]
		getAccounts = append(getAccounts, s.domainAccountToGetAccount(domainAccount))
	}

	return getAccounts
}
