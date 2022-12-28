package deposit

type Deposit struct {
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	PayeeAccountId  uint
	Amount          int
	Type            string
}
