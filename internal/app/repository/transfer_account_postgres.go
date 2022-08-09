package repository

import (
	"errors"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"gorm.io/gorm"
)

type TransferAccountPostgres struct {
	db *gorm.DB
}

func NewTransferAccountPostgres(db *gorm.DB) *TransferAccountPostgres {
	return &TransferAccountPostgres{db: db}
}

func (r *TransferAccountPostgres) GetSenderAccount(userId uint, accountId uint) (*domain.Account, error) {
	account := domain.NewAccount()

	result := r.db.Table("accounts").
		Where("user_id = ?", userId).
		Where("id = ?", accountId).
		First(&account)

	if result.Error != nil {
		customError := errors.New(responses.SenderAccountNotFound)
		return account, customError
	}

	return account, nil
}

func (r *TransferAccountPostgres) GetClientPayeeAccount(accountId uint) (*domain.Account, error) {
	account := domain.NewAccount()

	result := r.db.Table("accounts").
		Where("id = ?", accountId).
		First(&account)

	if result.Error != nil {
		customError := errors.New(responses.PayeeAccountNotFound)
		return account, customError
	}

	return account, nil
}

func (r *TransferAccountPostgres) GetPersonalPayeeAccount(userId uint, accountId uint) (*domain.Account, error) {
	account := domain.NewAccount()

	result := r.db.Table("accounts").
		Where("user_id = ?", userId).
		Where("id = ?", accountId).
		First(&account)

	if result.Error != nil {
		customError := errors.New(responses.PayeeAccountNotFound)
		return account, customError
	}

	return account, nil
}

func (r *TransferAccountPostgres) GetAccountDbObject(accountId uint) (*gorm.DB, error) {
	dbAccount := r.db.Table("accounts").
		Where("id = ?", accountId)

	if dbAccount.Error != nil {
		customError := errors.New(responses.AccountNotFound)
		return nil, customError
	}

	return dbAccount, nil
}

func (r *TransferAccountPostgres) UpdateAccountBalanceDbObject(dbAccountObject *gorm.DB, balance int) error {
	err := dbAccountObject.Update("balance", balance)
	if err.Error != nil {
		customError := errors.New(responses.CannotMakeDatabaseUpdate)
		return customError
	}
	return nil
}
