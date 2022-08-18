package balanceserv

import (
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"github.com/vaberof/MockBankingApplication/internal/app/repository"
)

type BalanceService struct {
	repos repository.Balance
}

func NewBalanceService(repos repository.Balance) *BalanceService {
	return &BalanceService{repos: repos}
}

func (s *BalanceService) GetBalance(userId uint) (*domain.Accounts, error) {
	accounts, err := s.repos.GetBalance(userId)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
