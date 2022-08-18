package userpg

import (
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"gorm.io/gorm"
)

type UserFinderPostgres struct {
	db *gorm.DB
}

func NewUserFinderPostgres(db *gorm.DB) *UserFinderPostgres {
	return &UserFinderPostgres{db: db}
}

func (r *UserFinderPostgres) GetUserById(userId uint) (*domain.User, error) {
	var user *domain.User

	result := r.db.Table("users").Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *UserFinderPostgres) GetUserByUsername(username string) (*domain.User, error) {
	var user *domain.User

	result := r.db.Table("users").Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
