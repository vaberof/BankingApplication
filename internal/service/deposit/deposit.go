package deposit

type Deposit struct {
	Id              uint
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	//PayeeUsername   string
	PayeeAccountId uint
	Amount         int
	Type           string
}
