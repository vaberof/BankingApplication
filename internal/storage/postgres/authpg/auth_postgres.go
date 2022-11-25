package authpg

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *domain.User) error {
	err := r.db.Create(&user)
	return err.Error
}
