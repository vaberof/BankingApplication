package accountserv

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/app/repository"
)

type AccountFinderService struct {
	repos repository.AccountFinder
}

func NewAccountFinderService(repos repository.AccountFinder) *AccountFinderService {
	return &AccountFinderService{repos: repos}
}

func (s *AccountFinderService) GetAccountByType(userId uint, accountType string) (*domain.Account, error) {
	account, err := s.repos.GetAccountByType(userId, accountType)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AccountFinderService) GetAccountById(accountId uint) (*domain.Account, error) {
	account, err := s.repos.GetAccountById(accountId)
	if err != nil {
		return nil, err
	}
	return account, nil

}
