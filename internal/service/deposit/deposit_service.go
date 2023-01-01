package deposit

import (
	"errors"
	"fmt"
)

type DepositService struct {
	depositStorage DepositStorage
}

func NewDepositService(depositStorage DepositStorage) *DepositService {
	return &DepositService{
		depositStorage: depositStorage,
	}
}

func (s *DepositService) SaveDeposit(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint) error {

	return s.depositStorage.SaveDeposit(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount)
}

func (s *DepositService) GetDeposits(userId uint) ([]*Deposit, error) {
	return s.getDepositsImpl(userId)
}

func (s *DepositService) getDepositsImpl(userId uint) ([]*Deposit, error) {
	deposits, err := s.depositStorage.GetDeposits(userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get deposits: %s", err.Error())
	}

	if len(deposits) == 0 {
		return nil, errors.New("there are no deposits")
	}

	return deposits, nil
}
