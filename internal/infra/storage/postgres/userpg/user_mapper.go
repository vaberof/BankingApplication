package userpg

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/user"
)

func BuildDomainUser(postgresUser *PostgresUser) *domain.User {
	return &domain.User{
		Id:       postgresUser.Id,
		Username: postgresUser.Username,
		Password: postgresUser.Password,
	}
}
