package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password []byte `json:"-"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(inputUsername string) {
	u.Username = inputUsername
}

func (u *User) SetPassword(inputPassword []byte) {
	u.Password = inputPassword
}
