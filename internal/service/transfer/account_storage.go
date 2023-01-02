package transfer

import "github.com/vaberof/MockBankingApplication/internal/domain/account"

type AccountStorage interface {
	GetAccountById(accountId uint) (*account.Account, error)
}
