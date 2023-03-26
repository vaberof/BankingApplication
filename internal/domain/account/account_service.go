package account

import (
	"errors"
	"fmt"
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

func (s *AccountService) CreateCustomAccount(userId uint, accountName string) (*Account, error) {
	return s.createCustomAccountImpl(userId, accountName)
}

func (s *AccountService) DeleteAccount(userId uint, accountName string) error {
	return s.deleteAccountImpl(userId, accountName)
}

func (s *AccountService) GetAccountByName(userId uint, accountName string) (*Account, error) {
	return s.getAccountByNameImpl(userId, accountName)
}

func (s *AccountService) GetAccountById(accountId uint) (*Account, error) {
	return s.getAccountByIdImpl(accountId)
}

func (s *AccountService) GetAccounts(userId uint) ([]*Account, error) {
	return s.getAccountsImpl(userId)
}

func (s *AccountService) createCustomAccountImpl(userId uint, accountName string) (*Account, error) {
	_, err := s.GetAccountByName(userId, accountName)
	if err == nil {
		return nil, errors.New("account with this name already exist")
	}

	account, err := s.accountStorage.CreateCustomAccount(userId, accountName)
	if err != nil {
		return nil, fmt.Errorf("cannot create account: %s", err.Error())
	}

	return account, nil
}

func (s *AccountService) deleteAccountImpl(userId uint, accountName string) error {
	account, err := s.accountStorage.GetAccountByName(userId, accountName)
	if err != nil {
		return err
	}

	if s.isMainAccountType(account.Type) {
		return errors.New("cannot delete account with main type")
	}

	if !s.isZeroAccountBalance(account.Balance) {
		return errors.New("cannot delete account with non-zero balance")
	}

	err = s.accountStorage.DeleteAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) getAccountByNameImpl(userId uint, accountName string) (*Account, error) {
	account, err := s.accountStorage.GetAccountByName(userId, accountName)
	if err != nil {
		return nil, errors.New("cannot find account with this name")
	}

	return account, nil
}

func (s *AccountService) getAccountByIdImpl(accountId uint) (*Account, error) {
	account, err := s.accountStorage.GetAccountById(accountId)
	if err != nil {
		return nil, errors.New("cannot find account with this id")
	}

	return account, nil
}

func (s *AccountService) getAccountsImpl(userId uint) ([]*Account, error) {
	accounts, err := s.accountStorage.GetAccounts(userId)
	if err != nil {
		return nil, errors.New("cannot find any accounts")
	}

	return accounts, nil
}

func (s *AccountService) isZeroAccountBalance(balance int) bool {
	return balance == 0
}

func (s *AccountService) isMainAccountType(accountType string) bool {
	return accountType == "Main"
}
