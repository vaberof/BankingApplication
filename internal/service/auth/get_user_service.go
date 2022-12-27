package auth

import (
	getuser "github.com/vaberof/MockBankingApplication/internal/service/user"
)

type GetUserService interface {
	GetUser(username string, password string) (*getuser.GetUserResponse, error)
	GetUserById(userId uint) (*getuser.GetUserResponse, error)
}
