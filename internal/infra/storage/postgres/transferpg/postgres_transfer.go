package transferpg

import "time"

type Transfer struct {
	Id              uint
	SenderId        uint
	SenderUsername  string
	SenderAccountId uint
	PayeeId         uint
	PayeeUsername   string
	PayeeAccountId  uint
	Amount          uint
	TransferType    string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
