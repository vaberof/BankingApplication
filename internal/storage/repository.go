package storage

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	accountpg2 "github.com/vaberof/banking_app/internal/storage/postgres/accountpg"
	authpg2 "github.com/vaberof/banking_app/internal/storage/postgres/authpg"
	"github.com/vaberof/banking_app/internal/storage/postgres/balancepg"
	"github.com/vaberof/banking_app/internal/storage/postgres/depositpg"
	transferpg2 "github.com/vaberof/banking_app/internal/storage/postgres/transferpg"
	"github.com/vaberof/banking_app/internal/storage/postgres/userpg"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user *domain.User) error
}

type AuthorizationValidator interface {
	UserExists(username string) error
}

type UserFinder interface {
	UserFinderById
	UserFinderByUsername
}

type UserFinderById interface {
	GetUserById(userId uint) (*domain.User, error)
}

type UserFinderByUsername interface {
	GetUserByUsername(username string) (*domain.User, error)
}

type Account interface {
	CreateAccount(account *domain.Account) error
	DeleteAccount(account *domain.Account) error
}

type AccountValidator interface {
	AccountExists(userId uint, accountType string) error
}

type AccountFinder interface {
	FindAccountById
	FindAccountByType
}

type FindAccountById interface {
	GetAccountById(accountId uint) (*domain.Account, error)
}

type FindAccountByType interface {
	GetAccountByType(userId uint, accountType string) (*domain.Account, error)
}

type Balance interface {
	GetBalance(userId uint) (*domain.Accounts, error)
}

type Transfer interface {
	CreateTransfer(transfer *domain.Transfer) error
	GetTransfers(userId uint) (*domain.Transfers, error)
}

type TransferAccount interface {
	GetSenderAccount(userId uint, accountId uint) (*domain.Account, error)
	GetClientPayeeAccount(accountId uint) (*domain.Account, error)
	GetPersonalPayeeAccount(userId uint, accountId uint) (*domain.Account, error)
	GetAccountDbObject(accountId uint) (*gorm.DB, error)
	UpdateAccountBalanceDbObject(dbAccountObject *gorm.DB, balance int) error
}

type Deposit interface {
	CreateDeposit(deposit *domain.Deposit) error
	GetDeposits(userId uint) (*domain.Deposits, error)
}

type Repository struct {
	Authorization
	AuthorizationValidator
	UserFinder
	Account
	AccountValidator
	AccountFinder
	Balance
	Transfer
	TransferAccount
	Deposit
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization:          authpg2.NewAuthPostgres(db),
		AuthorizationValidator: authpg2.NewAuthValidationPostgres(db),
		UserFinder:             userpg.NewUserFinderPostgres(db),
		Account:                accountpg2.NewAccountPostgres(db),
		AccountValidator:       accountpg2.NewAccountValidationPostgres(db),
		AccountFinder:          accountpg2.NewAccountFinderPostgres(db),
		Balance:                balancepg.NewBalancePostgres(db),
		Transfer:               transferpg2.NewTransferPostgres(db),
		TransferAccount:        transferpg2.NewTransferAccountPostgres(db),
		Deposit:                depositpg.NewDepositPostgres(db),
	}
}

func (r *Repository) MakeMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(domain.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(domain.Account{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(domain.Transfer{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(domain.Deposit{})
	if err != nil {
		return err
	}

	return nil
}
