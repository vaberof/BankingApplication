package account

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/account"
)

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccount(userId uint, accountName string) (*domain.Account, error)
	GetAccounts(userId uint) ([]*domain.Account, error)
	DeleteAccount(userId uint, accountName string) error
	UpdateBalance(userId uint, accountName string, balance int) error
}
