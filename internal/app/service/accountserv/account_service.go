package accountserv

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/storage"
)

type AccountService struct {
	repos storage.Account
}

func NewAccountService(repos storage.Account) *AccountService {
	return &AccountService{repos: repos}
}

func (s *AccountService) CreateInitialAccount(userId uint, username string) error {
	account := domain.NewAccount()

	account.SetUserId(userId)
	account.SetOwner(username)
	account.SetInitialMainAccountType()
	account.SetInitialBalance()

	return s.repos.CreateAccount(account)
}

func (s *AccountService) CreateCustomAccount(userId uint, username string, accountType string) error {
	account := domain.NewAccount()

	account.SetUserId(userId)
	account.SetOwner(username)
	account.SetCustomAccountType(accountType)

	return s.repos.CreateAccount(account)
}

func (s *AccountService) DeleteAccount(account *domain.Account) error {
	return s.repos.DeleteAccount(account)
}
