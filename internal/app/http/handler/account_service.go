package handler

import (
	"github.com/vaberof/MockBankingApplication/internal/service/account"
)

type AccountService interface {
	CreateCustomAccount(userId uint, accountName string) error
	GetAccounts(userId uint) ([]*account.GetAccountResponse, error)
	DeleteAccount(userId uint, accountName string) error
}
