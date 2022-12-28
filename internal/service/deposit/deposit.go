package deposit

import "time"

type Deposit struct {
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	PayeeAccountId  uint
	Amount          uint
	Date            time.Time
}
