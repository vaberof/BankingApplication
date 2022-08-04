package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/domain/deposit"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func CreateDeposit(senderUserID uint, senderUsername string, senderAccountID, payeeID, payeeAccountID uint, amount int, transferType string) *deposit.Deposit {
	newDeposit := deposit.NewDeposit()

	newDeposit.SetSenderID(senderUserID)
	newDeposit.SetSenderUsername(senderUsername)
	newDeposit.SetSenderAccountID(senderAccountID)
	newDeposit.SetPayeeID(payeeID)
	newDeposit.SetPayeeAccountID(payeeAccountID)
	newDeposit.SetAmount(amount)
	newDeposit.SetType(transferType)

	return newDeposit
}

func CreateDepositInDatabase(deposit *deposit.Deposit) {
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

func GetUserDeposits(claims *jwt.RegisteredClaims) (*deposit.Deposits, error) {
	var deposits *deposit.Deposits

	database.DB.Table("deposits").Where("payee_id = ?", claims.Issuer).Find(&deposits)

	dereferenceDeposits := *deposits

	if len(dereferenceDeposits) == 0 {
		customError := errors.New(responses.DepositsNotFound)
		return deposits, customError
	}

	return deposits, nil
}
