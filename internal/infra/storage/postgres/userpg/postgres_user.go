package userpg

import "time"

type User struct {
	Id        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
