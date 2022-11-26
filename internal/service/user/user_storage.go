package user

import (
	domain "github.com/vaberof/banking_app/internal/domain/user"
)

type UserStorage interface {
	CreateUser(username string, password string) error
	GetUser(username string) (*domain.User, error)
	GetUserById(userId uint) (*domain.User, error)
}
