package account

import "errors"

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
	return s.createCustomAccountImpl(userId, accountName)
}

func (s *AccountService) DeleteAccount(userId uint, accountName string) error {
	return s.accountStorage.DeleteAccount(userId, accountName)
}

func (s *AccountService) GetAccountByName(userId uint, accountName string) (*Account, error) {
	return s.getAccountByNameImpl(userId, accountName)
}

func (s *AccountService) GetAccountById(userId uint, accountId uint) (*Account, error) {
	return s.getAccountByIdImpl(userId, accountId)
}

func (s *AccountService) GetAccounts(userId uint) ([]*Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *AccountService) createCustomAccountImpl(userId uint, accountName string) error {
	account, err := s.GetAccountByName(userId, accountName)
	if account != nil || err == nil {
		return errors.New("account with this name already exist")
	}

	err = s.accountStorage.CreateCustomAccount(userId, accountName)
	if err != nil {
		return errors.New("cannot create account")
	}

	return err
}

func (s *AccountService) getAccountByNameImpl(userId uint, accountName string) (*Account, error) {
	account, err := s.accountStorage.GetAccountByName(userId, accountName)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) getAccountByIdImpl(userId uint, accountId uint) (*Account, error) {
	account, err := s.accountStorage.GetAccountById(userId, accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (s *AccountService) getAccountsImpl(userId uint) ([]*Account, error) {
	accounts, err := s.accountStorage.GetAccounts(userId)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
