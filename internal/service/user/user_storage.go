package user

import (
	domain "github.com/vaberof/banking_app/internal/domain/user"
)

type UserStorage interface {
	CreateUser(username string, password string) error
	GetUserById(userId uint) (*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}
