package transfer

import "github.com/vaberof/MockBankingApplication/internal/service/user"

type UserResponseService interface {
	GetUserById(userId uint) (*user.GetUserResponse, error)
}
