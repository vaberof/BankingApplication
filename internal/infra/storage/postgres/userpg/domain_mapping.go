package userpg

import domain "github.com/vaberof/MockBankingApplication/internal/domain/user"

func (s *PostgresUserStorage) infraUserToDomain(infraUser *User) *domain.User {
	var user domain.User

	user.Id = infraUser.Id
	user.Username = infraUser.Username
	user.Password = infraUser.Password

	return &user
}
