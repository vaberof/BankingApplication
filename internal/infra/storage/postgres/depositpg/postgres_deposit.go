package depositpg

import "time"

type Deposit struct {
	Id              uint
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	PayeeAccountId  uint
	Amount          uint
	DepositType     string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
