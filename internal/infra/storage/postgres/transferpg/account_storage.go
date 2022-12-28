package transferpg

import (
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/accountpg"
)

type AccountStorage interface {
	GetAccountById(accountId uint) (*account.Account, error)
	UpdateBalance(account *accountpg.Account, balance int) error
	DomainAccountToInfra(domainAccount *account.Account) *accountpg.Account
}
