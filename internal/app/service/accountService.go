package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/model"
)

func CreateAccount(accountType string, claims *jwt.RegisteredClaims) *model.Account {
	account := model.NewAccount()
	user := model.NewUser()

	userID := claims.Issuer
	database.DB.Where("id = ?", userID).First(&user)

	account.SetUserID(user.ID)
	account.SetOwner(user.Username)
	account.SetCustomAccountType(accountType)

	return account
}

func CreateAccountInDatabase(account *model.Account) {
	database.DB.Create(&account)
}

func DeleteAccountFromDatabase(account *model.Account) {
	database.DB.Delete(&account)
}

func GetUserAccounts(claims *jwt.RegisteredClaims) (*model.Accounts, error) {
	var accounts *model.Accounts

	database.DB.Table("accounts").Where("user_id = ?", claims.Issuer).Find(&accounts)

	dereferenceAccounts := *accounts

	if len(*dereferenceAccounts) == 0 {
		customError := errors.New(constants.TransactionsNotFound)
		return accounts, customError
	}

	return accounts, nil
}

func FindAccountByID(accountID string) (*model.Account, error) {
	account := model.NewAccount()

	result := database.DB.Table("accounts").Where("id = ?", accountID).First(&account)
	if result.Error != nil {
		return account, result.Error
	}

	return account, nil
}

func FindAccountByType(accountType string, claims *jwt.RegisteredClaims) (*model.Account, error) {
	account := model.NewAccount()

	userID := claims.Issuer

	result := database.DB.Table("accounts").Where("user_id = ?", userID).Where("type = ?", accountType).First(&account)
	if result.Error != nil {
		return account, result.Error
	}

	return account, nil
}

func IsMainAccountType(accountType string) bool {
	return accountType == "Main"
}

func IsEmptyAccountType(accountType string) bool {
	return accountType == ""
}

func IsZeroBalance(account *model.Account) bool {
	return account.Balance == 0
}
