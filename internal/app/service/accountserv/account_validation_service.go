package accountserv

import (
	"github.com/vaberof/banking_app/internal/storage"
)

type AccountValidationService struct {
	repos storage.AccountValidator
}

func NewAccountValidationService(repos storage.AccountValidator) *AccountValidationService {
	return &AccountValidationService{repos: repos}
}

func (s *AccountValidationService) AccountExists(userId uint, accountType string) error {
	return s.repos.AccountExists(userId, accountType)
}

func (s *AccountValidationService) IsEmptyAccountType(accountType string) bool {
	return len(accountType) == 0
}

func (s *AccountValidationService) IsMainAccountType(accountType string) bool {
	return accountType == "Main"
}

func (s *AccountValidationService) IsZeroBalance(balance int) bool {
	return balance == 0
}
