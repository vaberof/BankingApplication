package domain

type Accounts []*Account

type Account struct {
	Id      uint   `json:"id" gorm:"primary"`
	UserId  uint   `json:"-"`
	Owner   string `json:"-"`
	Type    string `json:"type"`
	Balance int    `json:"balance"`
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
