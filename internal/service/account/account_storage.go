package account

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/account"
)

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccountByName(userId uint, accountName string) (*domain.Account, error)
	GetAccountById(accountId uint) (*domain.Account, error)
	GetAccounts(userId uint) ([]*domain.Account, error)
	DeleteAccount(userId uint, accountName string) error
}
