package model

import "gorm.io/gorm"

type Transfers *[]Transfer

type Transfer struct {
	gorm.Model

	SenderID        uint
	SenderAccountID uint

	PayeeUsername  string
	PayeeAccountID uint

	Amount int
	Type   string
}

func NewTransfer() *Transfer {
	return &Transfer{}
}

func (t *Transfer) SetSenderID(senderID uint) {
	t.SenderID = senderID
}

func (t *Transfer) SetSenderAccountID(senderAccountID uint) {
	t.SenderAccountID = senderAccountID
}

func (t *Transfer) SetPayeeUsername(payeeUsername string) {
	t.PayeeUsername = payeeUsername
}

func (t *Transfer) SetPayeeAccountID(payeeAccountId uint) {
	t.PayeeAccountID = payeeAccountId
}

func (t *Transfer) SetAmount(amount int) {
	t.Amount = amount
}

func (t *Transfer) SetType(transferType string) {
	t.Type = transferType
}
