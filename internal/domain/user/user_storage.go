package user

import (
	infra "github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/userpg"
)

type UserStorage interface {
	CreateUser(username string, password string) error
	GetUserById(userId uint) (*infra.User, error)
	GetUserByUsername(username string) (*infra.User, error)
}
