package transferpg

import "time"

type Transfer struct {
	Id              uint
	SenderId        uint
	SenderAccountId uint
	PayeeId         uint
	PayeeAccountId  uint
	Amount          uint
	TransferType    string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
