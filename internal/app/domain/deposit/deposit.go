package deposit

import "gorm.io/gorm"

type Deposits []*Deposit

type Deposit struct {
	gorm.Model

	SenderID        uint
	SenderUsername  string
	SenderAccountID uint

	PayeeID        uint
	PayeeAccountID uint

	Amount int
	Type   string
}

func NewDeposit() *Deposit {
	return &Deposit{}
}

func (d *Deposit) SetSenderID(senderID uint) {
	d.SenderID = senderID
}

func (d *Deposit) SetSenderUsername(username string) {
	d.SenderUsername = username
}

func (d *Deposit) SetSenderAccountID(senderAccountID uint) {
	d.SenderAccountID = senderAccountID
}

func (d *Deposit) SetPayeeID(payeeID uint) {
	d.PayeeID = payeeID
}

func (d *Deposit) SetPayeeAccountID(payeeAccountId uint) {
	d.PayeeAccountID = payeeAccountId
}

func (d *Deposit) SetAmount(amount int) {
	d.Amount = amount
}

func (d *Deposit) SetType(transferType string) {
	d.Type = transferType
}
