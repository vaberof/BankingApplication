package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"github.com/vaberof/MockBankingApplication/internal/app/repository"
	"github.com/vaberof/MockBankingApplication/internal/app/service/accountserv"
	"github.com/vaberof/MockBankingApplication/internal/app/service/authserv"
	"github.com/vaberof/MockBankingApplication/internal/app/service/balanceserv"
	"github.com/vaberof/MockBankingApplication/internal/app/service/depositserv"
	"github.com/vaberof/MockBankingApplication/internal/app/service/transferserv"
	"github.com/vaberof/MockBankingApplication/internal/app/service/userserv"
)

type Authorization interface {
	CreateUser(user *domain.User) error
	GenerateJwtToken(userId uint) (string, error)
	ParseJwtToken(token string) (*jwt.Token, error)
	GenerateCookie(token string) *fiber.Cookie
	RemoveCookie() *fiber.Cookie
}

type AuthorizationValidator interface {
	UserExists(username string) error
	IsCorrectPassword(hashedPassword string, inputPassword string) error
	IsEmptyUsername(username string) bool
	IsEmptyPassword(password string) bool
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
	CreateInitialAccount(userId uint, username string) error
	CreateCustomAccount(userId uint, username string, accountType string) error
	DeleteAccount(account *domain.Account) error
}

type AccountValidator interface {
	AccountExists(userId uint, accountType string) error
	IsEmptyAccountType(accountType string) bool
	IsMainAccountType(accountType string) bool
	IsZeroBalance(balance int) bool
}

type AccountFinder interface {
	FindAccountByType
}

type FindAccountByType interface {
	GetAccountByType(userId uint, accountType string) (*domain.Account, error)
}

type Balance interface {
	GetBalance(userId uint) (*domain.Accounts, error)
}

type Transfer interface {
	CreateTransfer(transfer *domain.Transfer) error
	MakeTransfer(transfer *domain.Transfer) error
	TransformInputToTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount int, transferType string) (*domain.Transfer, error)
	GetTransfers(userId uint) (*domain.Transfers, error)
}

type Deposit interface {
	CreateDeposit(deposit *domain.Deposit) error
	ConvertTransferToDeposit(transfer *domain.Transfer) (*domain.Deposit, error)
	GetDeposits(userId uint) (*domain.Deposits, error)
}

type Service struct {
	Authorization
	AuthorizationValidator
	UserFinder
	Account
	AccountValidator
	AccountFinder
	Balance
	Transfer
	Deposit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:          authserv.NewAuthService(repos.Authorization),
		AuthorizationValidator: authserv.NewAuthValidationService(repos.AuthorizationValidator),
		UserFinder:             userserv.NewUserFinderService(repos.UserFinder),
		Account:                accountserv.NewAccountService(repos.Account),
		AccountValidator:       accountserv.NewAccountValidationService(repos.AccountValidator),
		AccountFinder:          accountserv.NewAccountFinderService(repos.AccountFinder),
		Balance:                balanceserv.NewBalanceService(repos.Balance),
		Transfer:               transferserv.NewTransferService(repos.Transfer, repos.TransferAccount, repos.AccountFinder),
		Deposit:                depositserv.NewDepositService(repos.Deposit, repos.UserFinder, repos.AccountFinder),
	}
}
