package handler

import (
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
)

type AccountService interface {
	CreateCustomAccount(userId uint, accountName string) (*account.Account, error)
	GetAccounts(userId uint) ([]*account.Account, error)
	DeleteAccount(userId uint, accountName string) error
}
