package domain

type Deposits []*Deposit

type Deposit struct {
	Id              uint   `json:"-" gorm:"primary"`
	SenderId        uint   `json:"-"`
	SenderUsername  string `json:"sender_username"`
	SenderAccountId uint   `json:"sender_account_id"`
	PayeeId         uint   `json:"-"`
	PayeeAccountId  uint   `json:"payee_account_id"`
	Amount          int    `json:"amount"`
	Type            string `json:"transfer_type"`
}

func NewDeposit() *Deposit {
	return &Deposit{}
}

func (d *Deposit) SetSenderId(senderId uint) {
	d.SenderId = senderId
}

func (d *Deposit) SetSenderUsername(username string) {
	d.SenderUsername = username
}

func (d *Deposit) SetSenderAccountId(senderAccountId uint) {
	d.SenderAccountId = senderAccountId
}

func (d *Deposit) SetPayeeId(payeeId uint) {
	d.PayeeId = payeeId
}

func (d *Deposit) SetPayeeAccountId(payeeAccountId uint) {
	d.PayeeAccountId = payeeAccountId
}

func (d *Deposit) SetAmount(amount int) {
	d.Amount = amount
}

func (d *Deposit) SetType(transferType string) {
	d.Type = transferType
}
