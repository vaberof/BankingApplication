package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/domain/account"
	"github.com/vaberof/banking_app/internal/app/domain/user"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func CreateCustomAccount(accountType string, claims *jwt.RegisteredClaims) *account.Account {
	newAccount := account.NewAccount()
	newUser := user.NewUser()

	userID := claims.Issuer
	database.DB.Where("id = ?", userID).First(&newUser)

	newAccount.SetUserID(newUser.ID)
	newAccount.SetOwner(newUser.Username)
	newAccount.SetCustomAccountType(accountType)

	return newAccount
}

func CreateInitialAccount(userID uint, username string) *account.Account {
	initialAccount := account.NewAccount()

	initialAccount.SetUserID(userID)
	initialAccount.SetOwner(username)
	initialAccount.SetMainAccountType()
	initialAccount.SetInitialBalance()

	return initialAccount
}

func CreateAccountInDatabase(account *account.Account) {
	database.DB.Create(&account)
}

func DeleteAccountFromDatabase(account *account.Account) {
	database.DB.Delete(&account)
}

func GetUserAccounts(claims *jwt.RegisteredClaims) (*account.Accounts, error) {
	var accounts *account.Accounts

	database.DB.Table("accounts").Where("user_id = ?", claims.Issuer).Find(&accounts)

	dereferenceAccounts := *accounts

	if len(dereferenceAccounts) == 0 {
		customError := errors.New(responses.TransfersNotFound)
		return accounts, customError
	}

	return accounts, nil
}

func FindAccountByID(accountID string) (*account.Account, error) {
	newAccount := account.NewAccount()

	result := database.DB.Table("accounts").Where("id = ?", accountID).First(&newAccount)
	if result.Error != nil {
		return newAccount, result.Error
	}

	return newAccount, nil
}

func FindAccountByType(accountType string, claims *jwt.RegisteredClaims) (*account.Account, error) {
	newAccount := account.NewAccount()

	userID := claims.Issuer

	result := database.DB.Table("accounts").Where("user_id = ?", userID).Where("type = ?", accountType).First(&newAccount)
	if result.Error != nil {
		return newAccount, result.Error
	}

	return newAccount, nil
}

func IsMainAccountType(accountType string) bool {
	return accountType == "Main"
}

func IsEmptyAccountType(accountType string) bool {
	return accountType == ""
}

func IsZeroBalance(account *account.Account) bool {
	return account.Balance == 0
}
