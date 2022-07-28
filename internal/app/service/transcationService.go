package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/model"
)

func CreateTransaction(senderUserID uint, payeeUsername string, payeeAccountID uint, amount int, transferType string) *model.Transaction {
	transaction := model.NewTransaction()

	transaction.SetUserID(senderUserID)
	transaction.SetPayeeUsername(payeeUsername)
	transaction.SetPayeeAccountID(payeeAccountID)
	transaction.SetAmount(amount)
	transaction.SetType(transferType)

	return transaction
}

func CreateTransactionInDatabase(transaction *model.Transaction) {
	database.DB.Create(&transaction)
}

func GetTransactionData(data map[string]string, claims *jwt.RegisteredClaims) (uint, string, uint, int, string) {
	customer, _ := FindUserById(claims.Issuer)
	payeeAccount, _ := FindAccountByID(data["to_account"])

	senderUserID := customer.ID
	payeeUsername := payeeAccount.Owner
	payeeAccountID := payeeAccount.ID
	amount := data["amount"]
	transferType := data["transfer_type"]
	intAmount, _ := ConvertAmountToInt(amount)

	return senderUserID, payeeUsername, payeeAccountID, intAmount, transferType
}

func GetUserTransactions(claims *jwt.RegisteredClaims) (*model.Transactions, error) {
	var transactions *model.Transactions

	database.DB.Table("transactions").Where("customer_id = ?", claims.Issuer).Find(&transactions)

	dereferenceTransactions := *transactions

	if len(*dereferenceTransactions) == 0 {
		customError := errors.New(constants.TransactionsNotFound)
		return transactions, customError
	}

	return transactions, nil
}
