package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/model"
	"gorm.io/gorm"
	"strconv"
)

func MakeTransfer(data map[string]string, claims *jwt.RegisteredClaims) error {
	transferType := data["transfer_type"]
	switch transferType {
	case "client":
		return ClientTransfer(data, claims)
	case "personal":
		return PersonalTransfer(data, claims)
	default:
		customError := errors.New(constants.UnsupportedTransferType)
		return customError
	}
}

func PersonalTransfer(data map[string]string, claims *jwt.RegisteredClaims) error {
	sentFromAccountID := data["from_account"]
	sentToAccountID := data["to_account"]
	amount := data["amount"]

	senderAccount, senderDbObject, err := getSenderAccount(claims, sentFromAccountID)
	if err != nil {
		return err
	}

	payeeAccount, payeeDbObject, err := getPersonalPayeeAccount(claims, sentToAccountID)
	if err != nil {
		return err
	}

	if isTheSameAccountID(senderAccount.ID, payeeAccount.ID) {
		customError := errors.New(constants.SenderIsRecipient)
		return customError
	}

	intAmount, err := ConvertAmountToInt(amount)
	if err != nil {
		return err
	}

	if !isEnoughFunds(senderAccount, intAmount) {
		customError := errors.New(constants.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderAccount.Balance - intAmount
	senderDbObject.Update("balance", newSenderBalance)

	newRecipientBalance := payeeAccount.Balance + intAmount
	payeeDbObject.Update("balance", newRecipientBalance)

	return nil
}

func ClientTransfer(data map[string]string, claims *jwt.RegisteredClaims) error {
	sentFromAccountID := data["from_account"]
	sentToAccountID := data["to_account"]
	amount := data["amount"]

	senderAccount, senderDbObject, err := getSenderAccount(claims, sentFromAccountID)
	if err != nil {
		return err
	}

	payeeAccount, payeeDbObject, err := getClientPayeeAccount(sentToAccountID)
	if err != nil {
		return err
	}

	if isTheSameAccountOwner(senderAccount.UserID, payeeAccount.UserID) {
		customError := errors.New(constants.SenderIsRecipient)
		return customError
	}

	intAmount, err := ConvertAmountToInt(amount)
	if err != nil {
		return err
	}

	if !isEnoughFunds(senderAccount, intAmount) {
		customError := errors.New(constants.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderAccount.Balance - intAmount
	senderDbObject.Update("balance", newSenderBalance)

	newRecipientBalance := payeeAccount.Balance + intAmount
	payeeDbObject.Update("balance", newRecipientBalance)

	return nil
}

func ConvertAmountToInt(amount string) (int, error) {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		customError := errors.New(constants.UnsupportedTransferAmount)
		return -1, customError
	}

	return intAmount, nil
}

func getSenderAccount(claims *jwt.RegisteredClaims, accountID string) (*model.Account, *gorm.DB, error) {
	account := model.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("user_id = ?", claims.Issuer).
		Where("id = ?", accountID).
		First(&account)

	if accountDbObject.Error != nil {
		customError := errors.New(constants.SenderAccountNotFound)
		return account, accountDbObject, customError
	}

	return account, accountDbObject, nil
}

func getPersonalPayeeAccount(claims *jwt.RegisteredClaims, accountID string) (*model.Account, *gorm.DB, error) {
	account := model.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("user_id = ?", claims.Issuer).
		Where("id = ?", accountID).
		First(&account)

	if accountDbObject.Error != nil {
		customError := errors.New(constants.PayeeAccountNotFound)
		return account, accountDbObject, customError
	}

	return account, accountDbObject, nil
}

func getClientPayeeAccount(accountID string) (*model.Account, *gorm.DB, error) {
	account := model.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("id = ?", accountID).
		First(&account)

	if accountDbObject.Error != nil {
		customError := errors.New(constants.PayeeAccountNotFound)
		return account, accountDbObject, customError
	}

	return account, accountDbObject, nil
}

func isTheSameAccountID(senderAccountID, recipientAccountID uint) bool {
	return senderAccountID == recipientAccountID
}

func isTheSameAccountOwner(senderUserID, recipientUserID uint) bool {
	return senderUserID == recipientUserID
}

func isEnoughFunds(account *model.Account, amount int) bool {
	return account.Balance-amount >= 0
}
