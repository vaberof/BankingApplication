package domain

type User struct {
	Id       uint   `json:"-" gorm:"primary"`
	Username string `json:"username"`
	Password string `json:"password"`
}
