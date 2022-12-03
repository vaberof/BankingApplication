package auth

import (
	getuser "github.com/vaberof/MockBankingApplication/internal/service/user"
)

type GetUserService interface {
	GetUser(username string, password string) (*getuser.GetUser, error)
	GetUserById(userId uint) (*getuser.GetUser, error)
}
