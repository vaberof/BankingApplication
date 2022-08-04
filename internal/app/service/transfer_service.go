package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/domain/account"
	"github.com/vaberof/banking_app/internal/app/domain/transfer"
	"github.com/vaberof/banking_app/internal/pkg/responses"
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
		customError := errors.New(responses.UnsupportedTransferType)
		return customError
	}
}

func PersonalTransfer(data map[string]string, claims *jwt.RegisteredClaims) error {
	senderAccountID := data["from_account"]
	payeeAccountID := data["to_account"]
	amount := data["amount"]

	senderAccount, senderDbObject, err := getSenderAccount(claims, senderAccountID)
	if err != nil {
		return err
	}

	payeeAccount, payeeDbObject, err := getPersonalPayeeAccount(claims, payeeAccountID)
	if err != nil {
		return err
	}

	if isTheSameAccountID(senderAccount.ID, payeeAccount.ID) {
		customError := errors.New(responses.SenderIsPayee)
		return customError
	}

	intAmount, err := ConvertAmountToInt(amount)
	if err != nil {
		return err
	}

	if !isEnoughFunds(senderAccount, intAmount) {
		customError := errors.New(responses.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderAccount.Balance - intAmount
	senderDbObject.Update("balance", newSenderBalance)

	newPayeeBalance := payeeAccount.Balance + intAmount
	payeeDbObject.Update("balance", newPayeeBalance)

	return nil
}

func ClientTransfer(data map[string]string, claims *jwt.RegisteredClaims) error {
	senderAccountID := data["from_account"]
	payeeAccountID := data["to_account"]
	amount := data["amount"]

	senderAccount, senderDbObject, err := getSenderAccount(claims, senderAccountID)
	if err != nil {
		return err
	}

	payeeAccount, payeeDbObject, err := getClientPayeeAccount(payeeAccountID)
	if err != nil {
		return err
	}

	if isTheSameAccountOwner(senderAccount.UserID, payeeAccount.UserID) {
		customError := errors.New(responses.SenderIsPayee)
		return customError
	}

	intAmount, err := ConvertAmountToInt(amount)
	if err != nil {
		return err
	}

	if !isEnoughFunds(senderAccount, intAmount) {
		customError := errors.New(responses.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderAccount.Balance - intAmount
	senderDbObject.Update("balance", newSenderBalance)

	newPayeeBalance := payeeAccount.Balance + intAmount
	payeeDbObject.Update("balance", newPayeeBalance)

	return nil
}

func CreateTransfer(senderUserID, senderAccountID uint, payeeUsername string, payeeAccountID uint, amount int, transferType string) *transfer.Transfer {
	newTransfer := transfer.NewTransfer()

	newTransfer.SetSenderID(senderUserID)
	newTransfer.SetSenderAccountID(senderAccountID)
	newTransfer.SetPayeeUsername(payeeUsername)
	newTransfer.SetPayeeAccountID(payeeAccountID)
	newTransfer.SetAmount(amount)
	newTransfer.SetType(transferType)

	return newTransfer
}

func CreateTransferInDatabase(transfer *transfer.Transfer) {
	database.DB.Create(&transfer)
}

func GetTransferData(data map[string]string, claims *jwt.RegisteredClaims) (uint, uint, string, uint, int, string) {
	sender, _ := FindUserById(claims.Issuer)
	senderAccount, _ := FindAccountByID(data["from_account"])
	payeeAccount, _ := FindAccountByID(data["to_account"])

	senderUserID := sender.ID
	senderAccountID := senderAccount.ID

	payeeUsername := payeeAccount.Owner
	payeeAccountID := payeeAccount.ID

	amount := data["amount"]
	intAmount, _ := ConvertAmountToInt(amount)

	transferType := data["transfer_type"]

	return senderUserID, senderAccountID, payeeUsername, payeeAccountID, intAmount, transferType
}

func GetUserTransfers(claims *jwt.RegisteredClaims) (*transfer.Transfers, error) {
	var transfers *transfer.Transfers

	database.DB.Table("transfers").Where("sender_id = ?", claims.Issuer).Find(&transfers)

	dereferenceTransfers := *transfers

	if len(dereferenceTransfers) == 0 {
		customError := errors.New(responses.TransfersNotFound)
		return transfers, customError
	}

	return transfers, nil
}

func ConvertAmountToInt(amount string) (int, error) {
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		customError := errors.New(responses.UnsupportedTransferAmount)
		return -1, customError
	}

	return intAmount, nil
}

func getSenderAccount(claims *jwt.RegisteredClaims, accountID string) (*account.Account, *gorm.DB, error) {
	newAccount := account.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("user_id = ?", claims.Issuer).
		Where("id = ?", accountID).
		First(&newAccount)

	if accountDbObject.Error != nil {
		customError := errors.New(responses.SenderAccountNotFound)
		return newAccount, accountDbObject, customError
	}

	return newAccount, accountDbObject, nil
}

func getPersonalPayeeAccount(claims *jwt.RegisteredClaims, accountID string) (*account.Account, *gorm.DB, error) {
	newAccount := account.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("user_id = ?", claims.Issuer).
		Where("id = ?", accountID).
		First(&newAccount)

	if accountDbObject.Error != nil {
		customError := errors.New(responses.PayeeAccountNotFound)
		return newAccount, accountDbObject, customError
	}

	return newAccount, accountDbObject, nil
}

func getClientPayeeAccount(accountID string) (*account.Account, *gorm.DB, error) {
	newAccount := account.NewAccount()

	accountDbObject := database.DB.Table("accounts").
		Where("id = ?", accountID).
		First(&newAccount)

	if accountDbObject.Error != nil {
		customError := errors.New(responses.PayeeAccountNotFound)
		return newAccount, accountDbObject, customError
	}

	return newAccount, accountDbObject, nil
}

func isTheSameAccountID(senderAccountID, payeeAccountID uint) bool {
	return senderAccountID == payeeAccountID
}

func isTheSameAccountOwner(senderUserID, payeeAccountID uint) bool {
	return senderUserID == payeeAccountID
}

func isEnoughFunds(account *account.Account, amount int) bool {
	return account.Balance-amount >= 0
}
