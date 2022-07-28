package model

import "gorm.io/gorm"

type Accounts *[]Account

type Account struct {
	gorm.Model
	UserID  uint
	Owner   string
	Type    string
	Balance int
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) SetUserID(userID uint) {
	a.UserID = userID
}

func (a *Account) SetOwner(username string) {
	a.Owner = username
}

func (a *Account) SetMainAccountType() {
	a.Type = "Main"
}

func (a *Account) SetCustomAccountType(accountType string) {
	a.Type = accountType
}

func (a *Account) SetInitialBalance() {
	a.Balance = 10000
}
