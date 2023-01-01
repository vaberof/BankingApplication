package authpg

import (
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"gorm.io/gorm"
)

type AuthorizationValidation struct {
	db *gorm.DB
}

func NewAuthValidationPostgres(db *gorm.DB) *AuthorizationValidation {
	return &AuthorizationValidation{db: db}
}

func (r *AuthorizationValidation) UserExists(username string) error {
	var user *domain.User

	result := r.db.Table("users").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
