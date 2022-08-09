package domain

type Accounts []*Account

type Account struct {
	Id      uint `gorm:"primary"`
	UserId  uint
	Owner   string
	Type    string
	Balance int
}

func NewAccount() *Account {
	return &Account{}
}

func (a *Account) SetUserId(userId uint) {
	a.UserId = userId
}

func (a *Account) SetOwner(username string) {
	a.Owner = username
}

func (a *Account) SetInitialMainAccountType() {
	a.Type = "Main"
}

func (a *Account) SetCustomAccountType(accountType string) {
	a.Type = accountType
}

func (a *Account) SetInitialBalance() {
	a.Balance = 10000
}
