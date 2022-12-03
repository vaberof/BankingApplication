package account

import (
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/accountpg"
)

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccount(userId uint, accountName string) (*accountpg.Account, error)
	GetAccounts(userId uint) ([]*accountpg.Account, error)
	DeleteAccount(userId uint, accountName string) error
	UpdateBalance(userId uint, accountName string, balance int) error
}
