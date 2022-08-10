package accountpg

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"gorm.io/gorm"
)

type AccountPostgres struct {
	db *gorm.DB
}

func NewAccountPostgres(db *gorm.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) CreateAccount(account *domain.Account) error {
	err := r.db.Create(&account)
	return err.Error
}

func (r *AccountPostgres) DeleteAccount(account *domain.Account) error {
	err := r.db.Delete(&account)
	return err.Error
}
