package depositpg

import "gorm.io/gorm"

type PostgresDepositStorage struct {
	db *gorm.DB
}

func NewPostgresDepositStorage(db *gorm.DB) *PostgresDepositStorage {
	return &PostgresDepositStorage{
		db: db,
	}
}

func (s *PostgresDepositStorage) SaveDeposit(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	depositType string) error {

	return s.saveDepositImpl(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount, depositType)
}

func (s *PostgresDepositStorage) saveDepositImpl(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	depositType string) error {

	var deposit Deposit

	deposit.SenderId = senderId
	deposit.SenderUsername = senderUsername
	deposit.SenderAccountId = senderAccountId
	deposit.PayeeId = payeeId
	deposit.PayeeUsername = payeeUsername
	deposit.PayeeAccountId = payeeAccountId
	deposit.Amount = amount
	deposit.DepositType = depositType

	err := s.db.Create(&deposit).Error
	if err != nil {
		return err
	}

	return nil
}
