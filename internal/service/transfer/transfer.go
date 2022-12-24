package transfer

type Transfer struct {
	SenderAccountId uint
	PayeeId         uint
	PayeeAccountId  uint
	Amount          uint
	TransferType    string
}
