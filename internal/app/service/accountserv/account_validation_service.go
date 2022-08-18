package accountserv

import (
	"github.com/vaberof/MockBankingApplication/internal/app/repository"
)

type AccountValidationService struct {
	repos repository.AccountValidator
}

func NewAccountValidationService(repos repository.AccountValidator) *AccountValidationService {
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
