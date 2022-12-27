package deposit

import domain "github.com/vaberof/MockBankingApplication/internal/domain/user"

type UserStorage interface {
	GetUserById(userId uint) (*domain.User, error)
}
