package handler

import "github.com/vaberof/banking_app/internal/service/user"

type UserService interface {
	CreateUser(username string, password string) error
	GetUserByUsername(username string) (*user.GetUser, error)
}
