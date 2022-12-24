package accountpg

import (
	"time"
)

type Account struct {
	Id        uint
	UserId    uint
	Type      string
	Name      string
	Balance   int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"autoDeleteTime" gorm:"index" `
}
