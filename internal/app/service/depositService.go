package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/model"
)

func CreateDeposit(senderUserID uint, senderUsername string, senderAccountID, payeeID, payeeAccountID uint, amount int, transferType string) *model.Deposit {
	deposit := model.NewDeposit()

	deposit.SetSenderID(senderUserID)
	deposit.SetSenderUsername(senderUsername)
	deposit.SetSenderAccountID(senderAccountID)
	deposit.SetPayeeID(payeeID)
	deposit.SetPayeeAccountID(payeeAccountID)
	deposit.SetAmount(amount)
	deposit.SetType(transferType)

	return deposit
}

func CreateDepositInDatabase(deposit *model.Deposit) {
	database.DB.Create(&deposit)
}

func GetDepositData(data map[string]string, claims *jwt.RegisteredClaims) (string, uint) {
	sender, _ := FindUserById(claims.Issuer)
	payeeAccount, _ := FindAccountByID(data["to_account"])
	payee, _ := FindUserByUsername(payeeAccount.Owner)

	senderUsername := sender.Username

	payeeID := payee.ID

	return senderUsername, payeeID
}

func GetUserDeposits(claims *jwt.RegisteredClaims) (*model.Deposits, error) {
	var deposits *model.Deposits

	database.DB.Table("deposits").Where("payee_id = ?", claims.Issuer).Find(&deposits)

	dereferenceTransfers := *deposits

	if len(*dereferenceTransfers) == 0 {
		customError := errors.New(constants.DepositsNotFound)
		return deposits, customError
	}

	return deposits, nil
}
