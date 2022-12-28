package transfer

type Transfer struct {
	SenderAccountId uint
	PayeeAccountId  uint
	PayeeUsername   string
	Amount          uint
	TransferType    string
}
