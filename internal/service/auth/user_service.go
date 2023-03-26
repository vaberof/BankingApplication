package auth

import (
	"github.com/vaberof/MockBankingApplication/internal/domain/user"
)

type UserService interface {
	GetUser(username string, password string) (*user.User, error)
	GetUserById(userId uint) (*user.User, error)
}
