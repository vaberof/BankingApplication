package repository

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"gorm.io/gorm"
)

type AccountFinderPostgres struct {
	db *gorm.DB
}

func NewAccountFinderPostgres(db *gorm.DB) *AccountFinderPostgres {
	return &AccountFinderPostgres{db: db}
}

func (r *AccountFinderPostgres) GetAccountByType(userId uint, accountType string) (*domain.Account, error) {
	var account *domain.Account

	result := r.db.Table("accounts").Where("user_id = ?", userId).Where("type = ?", accountType).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

func (r *AccountFinderPostgres) GetAccountById(accountId uint) (*domain.Account, error) {
	var account *domain.Account
	result := r.db.Table("accounts").Where("id = ?", accountId).First(&account)
	if result.Error != nil {
		return account, result.Error
	}

	return account, nil
}
