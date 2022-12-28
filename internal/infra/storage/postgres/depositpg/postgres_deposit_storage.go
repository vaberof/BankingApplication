package depositpg

import (
	"github.com/sirupsen/logrus"
	"github.com/vaberof/MockBankingApplication/internal/service/deposit"
	"gorm.io/gorm"
)

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
	amount uint) error {

	return s.saveDepositImpl(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount)
}

func (s *PostgresDepositStorage) GetDeposits(userId uint) ([]*deposit.Deposit, error) {
	return s.getDepositsImpl(userId)
}

func (s *PostgresDepositStorage) saveDepositImpl(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint) error {

	var deposit Deposit

	deposit.SenderId = senderId
	deposit.SenderUsername = senderUsername
	deposit.SenderAccountId = senderAccountId
	deposit.PayeeId = payeeId
	deposit.PayeeUsername = payeeUsername
	deposit.PayeeAccountId = payeeAccountId
	deposit.Amount = amount
	err := s.db.Create(&deposit).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "saveDepositImpl",
		}).Error(err)

		return err
	}

	return nil
}

func (s *PostgresDepositStorage) getDepositsImpl(userId uint) ([]*deposit.Deposit, error) {
	var infraDeposits []*Deposit

	err := s.db.Table("deposits").Where("payee_id = ?", userId).Find(&infraDeposits).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"layer": "infra",
			"func":  "getDepositsImpl",
		}).Error(err)

		return nil, err
	}

	serviceDeposits := s.infraDepositsToService(infraDeposits)

	return serviceDeposits, nil
}
