package account

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountType string, accountName string) error
	UpdateBalance(userId uint, accountName string, balance int) error
	DeleteAccount(userId uint, accountName string) error
}
