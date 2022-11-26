package accountpg

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId  uint
	Type    string
	Name    string
	Balance int
}
