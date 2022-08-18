package accountpg

import (
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"gorm.io/gorm"
)

type AccountValidationPostgres struct {
	db *gorm.DB
}

func NewAccountValidationPostgres(db *gorm.DB) *AccountValidationPostgres {
	return &AccountValidationPostgres{db: db}
}

func (r *AccountValidationPostgres) AccountExists(userId uint, accountType string) error {
	var account *domain.Account

	result := r.db.Table("accounts").Where("user_id = ?", userId).Where("type = ?", accountType).First(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
