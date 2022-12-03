package handler

import "github.com/vaberof/MockBankingApplication/internal/service/user"

type UserService interface {
	CreateUser(username string, password string) error
	GetUserByUsername(username string) (*user.GetUser, error)
}
