package transfer

import "time"

type Transfer struct {
	SenderAccountId uint
	PayeeAccountId  uint
	PayeeUsername   string
	Amount          uint
	TransferType    string
	Date            time.Time
}
