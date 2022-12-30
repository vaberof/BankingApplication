package accountpg

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/account"
)

func BuildDomainAccount(postgresAccount *PostgresAccount) *domain.Account {
	return &domain.Account{
		Id:      postgresAccount.Id,
		UserId:  postgresAccount.UserId,
		Type:    postgresAccount.Type,
		Name:    postgresAccount.Name,
		Balance: postgresAccount.Balance,
	}
}

func BuildDomainAccounts(postgresAccounts []*PostgresAccount) []*domain.Account {
	domainAccounts := make([]*domain.Account, len(postgresAccounts))

	for i := 0; i < len(domainAccounts); i++ {
		domainAccounts[i] = BuildDomainAccount(postgresAccounts[i])
	}

	return domainAccounts
}

func BuildPostgresAccount(domainAccount *domain.Account) *PostgresAccount {
	return &PostgresAccount{
		Id:      domainAccount.Id,
		UserId:  domainAccount.UserId,
		Type:    domainAccount.Type,
		Name:    domainAccount.Name,
		Balance: domainAccount.Balance,
	}
}
