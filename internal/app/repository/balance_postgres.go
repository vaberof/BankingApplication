package repository

import (
	"errors"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"gorm.io/gorm"
)

type BalancePostgres struct {
	db *gorm.DB
}

func NewBalancePostgres(db *gorm.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetBalance(userId uint) (*domain.Accounts, error) {
	var accounts *domain.Accounts

	r.db.Table("accounts").Where("user_id = ?", userId).Find(&accounts)

	if len(*accounts) == 0 {
		customError := errors.New(responses.TransfersNotFound)
		return accounts, customError
	}

	return accounts, nil
}
