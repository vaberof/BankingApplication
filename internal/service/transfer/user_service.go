package transfer

import "github.com/vaberof/MockBankingApplication/internal/domain/user"

type UserService interface {
	GetUserById(userId uint) (*user.User, error)
}
