package handler

import (
	"github.com/vaberof/MockBankingApplication/internal/service/account"
)

type AccountService interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccount(userId uint, accountName string) (*account.GetAccount, error)
	GetAccounts(userId uint) ([]*account.GetAccount, error)
	DeleteAccount(userId uint, accountName string) error
	UpdateBalance(userId uint, accountName string, balance int) error
}
