package accountpg

import "github.com/vaberof/MockBankingApplication/internal/domain/account"

func (s *PostgresAccountStorage) infraAccountToDomain(infraAccount *Account) *account.Account {
	var domainAccount account.Account

	domainAccount.Id = infraAccount.Id
	domainAccount.UserId = infraAccount.UserId
	domainAccount.Type = infraAccount.Type
	domainAccount.Name = infraAccount.Name
	domainAccount.Balance = infraAccount.Balance

	return &domainAccount
}

func (s *PostgresAccountStorage) infraAccountsToDomain(infraAccounts []*Account) []*account.Account {
	var domainAccounts []*account.Account

	for i := 0; i < len(infraAccounts); i++ {
		infraAccount := infraAccounts[i]
		domainAccounts = append(domainAccounts, s.infraAccountToDomain(infraAccount))
	}

	return domainAccounts
}

func (s *PostgresAccountStorage) DomainAccountToInfra(domainAccount *account.Account) *Account {
	var infraAccount Account

	infraAccount.Id = domainAccount.Id
	infraAccount.UserId = domainAccount.UserId
	infraAccount.Type = domainAccount.Type
	infraAccount.Name = domainAccount.Name
	infraAccount.Balance = domainAccount.Balance

	return &infraAccount
}

func (s *PostgresAccountStorage) DomainAccountsToInfra(domainAccounts []*account.Account) []*Account {
	var infraAccounts []*Account

	for i := 0; i < len(infraAccounts); i++ {
		infraAccount := infraAccounts[i]
		domainAccounts = append(domainAccounts, s.infraAccountToDomain(infraAccount))
	}

	return infraAccounts
}
