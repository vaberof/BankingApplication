package handler

import (
	"github.com/vaberof/banking_app/internal/service/account"
)

type AccountService interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccount(userId uint, accountName string) (*account.GetAccount, error)
	GetAccounts(userId uint) ([]*account.GetAccount, error)
	UpdateBalance(userId uint, accountName string, balance int) error
	DeleteAccount(userId uint, accountName string) error
}
