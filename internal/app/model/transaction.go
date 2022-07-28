package model

import "gorm.io/gorm"

type Transactions *[]Transaction

type Transaction struct {
	gorm.Model
	CustomerID     uint
	PayeeUsername  string
	PayeeAccountID uint
	Amount         int
	Type           string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) SetUserID(userID uint) {
	t.CustomerID = userID
}

func (t *Transaction) SetPayeeUsername(payeeUsername string) {
	t.PayeeUsername = payeeUsername
}

func (t *Transaction) SetPayeeAccountID(payeeAccountId uint) {
	t.PayeeAccountID = payeeAccountId
}

func (t *Transaction) SetAmount(amount int) {
	t.Amount = amount
}

func (t *Transaction) SetType(transferType string) {
	t.Type = transferType
}
