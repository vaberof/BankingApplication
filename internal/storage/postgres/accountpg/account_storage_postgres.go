package accountpg

import "gorm.io/gorm"

const (
	initialAccountType    = "Main"
	initialAccountName    = "General"
	initialAccountBalance = 10000

	SecondaryAccountType  = "Secondary"
	DefaultAccountBalance = 0
)

type AccountStoragePostgres struct {
	db *gorm.DB
}

func NewAccountStoragePostgres(db *gorm.DB) *AccountStoragePostgres {
	return &AccountStoragePostgres{db: db}
}

func (s *AccountStoragePostgres) CreateInitialAccount(userId uint) error {
	return s.createInitialAccountImpl(userId)
}

func (s *AccountStoragePostgres) CreateCustomAccount(userId uint, accountType string, accountName string) error {
	return s.createCustomAccountImpl(userId, accountType, accountName)
}

func (s *AccountStoragePostgres) UpdateBalance(userId uint, accountName string, balance int) error {
	return s.updateBalanceImpl(userId, accountName, balance)
}

func (s *AccountStoragePostgres) DeleteAccount(userId uint, accountName string) error {
	return s.deleteAccountImpl(userId, accountName)
}

func (s *AccountStoragePostgres) createInitialAccountImpl(userId uint) error {
	var account Account

	account.UserId = userId
	account.Type = initialAccountType
	account.Name = initialAccountName
	account.Balance = initialAccountBalance

	err := s.db.Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountStoragePostgres) createCustomAccountImpl(userId uint, accountType string, accountName string) error {
	var account Account

	account.UserId = userId
	account.Type = SecondaryAccountType
	account.Name = accountName
	account.Balance = DefaultAccountBalance

	err := s.db.Create(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountStoragePostgres) updateBalanceImpl(userId uint, accountName string, balance int) error {
	account, err := s.getAccount(userId, accountName)
	if err != nil {
		return err
	}

	account.Balance = balance
	err = s.db.Save(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountStoragePostgres) deleteAccountImpl(userId uint, accountName string) error {
	account, err := s.getAccount(userId, accountName)
	if err != nil {
		return err
	}

	err = s.db.Delete(&account).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountStoragePostgres) getAccount(userId uint, accountName string) (*Account, error) {
	var account Account

	err := s.db.Table("accounts").Where("user_id = ? AND account_name = ?", userId, accountName).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}
