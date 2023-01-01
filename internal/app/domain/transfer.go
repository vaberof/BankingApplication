package domain

type Transfers []*Transfer

type Transfer struct {
	Id              uint   `json:"-" gorm:"primary"`
	SenderId        uint   `json:"-"`
	SenderAccountId uint   `json:"sender_account_id"`
	PayeeUsername   string `json:"payee_username"`
	PayeeAccountId  uint   `json:"payee_account_id"`
	Amount          int    `json:"amount"`
	Type            string `json:"transfer_type"`
}

func NewTransfer() *Transfer {
	return &Transfer{}
}

func (t *Transfer) SetSenderId(senderId uint) {
	t.SenderId = senderId
}

func (t *Transfer) SetSenderAccountId(senderAccountId uint) {
	t.SenderAccountId = senderAccountId
}

func (t *Transfer) SetPayeeUsername(payeeUsername string) {
	t.PayeeUsername = payeeUsername
}

func (t *Transfer) SetPayeeAccountId(payeeAccountId uint) {
	t.PayeeAccountId = payeeAccountId
}

func (t *Transfer) SetAmount(amount int) {
	t.Amount = amount
}

func (t *Transfer) SetType(transferType string) {
	t.Type = transferType
}
