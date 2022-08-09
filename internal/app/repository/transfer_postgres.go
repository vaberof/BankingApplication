package repository

import (
	"errors"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"gorm.io/gorm"
)

type TransferPostgres struct {
	db *gorm.DB
}

func NewTransferPostgres(db *gorm.DB) *TransferPostgres {
	return &TransferPostgres{db: db}
}

func (r *TransferPostgres) CreateTransfer(transfer *domain.Transfer) error {
	err := r.db.Create(&transfer)
	return err.Error
}

func (r *TransferPostgres) GetTransfers(userId uint) (*domain.Transfers, error) {
	var transfers *domain.Transfers

	r.db.Table("transfers").Where("sender_id = ?", userId).Find(&transfers)

	if len(*transfers) == 0 {
		customError := errors.New(responses.TransfersNotFound)
		return transfers, customError
	}

	return transfers, nil
}
