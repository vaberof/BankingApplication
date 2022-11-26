package auth

import (
	getuser "github.com/vaberof/banking_app/internal/service/user"
)

type GetUserService interface {
	GetUser(username string, password string) (*getuser.GetUser, error)
	GetUserById(userId uint) (*getuser.GetUser, error)
}
