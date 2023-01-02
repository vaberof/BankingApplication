package handler

import "github.com/vaberof/MockBankingApplication/internal/domain/user"

type UserService interface {
	CreateUser(username string, password string) (*user.User, error)
}
