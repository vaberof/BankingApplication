package user

import (
	infra "github.com/vaberof/banking_app/internal/infra/storage/postgres/userpg"
)

type UserStorage interface {
	CreateUser(username string, password string) error
	GetUser(username string) (*infra.User, error)
	GetUserById(userId uint) (*infra.User, error)
}
