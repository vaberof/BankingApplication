package account

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
	CreateCustomAccount(userId uint, accountName string) error
	GetAccountByName(userId uint, accountName string) (*Account, error)
	GetAccountById(accountId uint) (*Account, error)
	GetAccounts(userId uint) ([]*Account, error)
	DeleteAccount(userId uint, accountName string) error
}
