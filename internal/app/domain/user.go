package domain

type User struct {
	Id       uint   `gorm:"primary"`
	Username string `json:"username"`
	Password string `json:"password"`
}
