package depositpg

import (
	"errors"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"gorm.io/gorm"
)

type DepositPostgres struct {
	db *gorm.DB
}

func (r *DepositPostgres) CreateDeposit(deposit *domain.Deposit) error {
	err := r.db.Create(&deposit)
	return err.Error
}

func NewDepositPostgres(db *gorm.DB) *DepositPostgres {
	return &DepositPostgres{db: db}
}

func (r *DepositPostgres) GetDeposits(userId uint) (*domain.Deposits, error) {
	var deposits *domain.Deposits

	r.db.Table("deposits").Where("payee_id = ?", userId).Find(&deposits)

	if len(*deposits) == 0 {
		customError := errors.New(responses.DepositsNotFound)
		return deposits, customError
	}

	return deposits, nil
}
