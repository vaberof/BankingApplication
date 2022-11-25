package account

type AccountService struct {
	accountStorage AccountStorage
}

func NewAccountService(accountStorage AccountStorage) *AccountService {
	return &AccountService{accountStorage: accountStorage}
}

func (s *AccountService) CreateInitialAccount(userId uint) error {
	return s.accountStorage.CreateInitialAccount(userId)
}

func (s *AccountService) CreateCustomAccount(userId uint, accountType string, accountName string) error {
	return s.accountStorage.CreateCustomAccount(userId, accountType, accountName)
}

func (s *AccountService) UpdateBalance(userId uint, accountName string, balance int) error {
	return s.accountStorage.UpdateBalance(userId, accountName, balance)
}

func (s *AccountService) DeleteAccount(userId uint, accountName string) error {
	return s.accountStorage.DeleteAccount(userId, accountName)
}
