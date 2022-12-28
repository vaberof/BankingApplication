package deposit

import (
	"errors"
	"github.com/vaberof/MockBankingApplication/internal/app/http/handler"
)

type DepositService struct {
	depositStorage DepositStorage
}

func NewDepositService(depositStorage DepositStorage) *DepositService {
	return &DepositService{
		depositStorage: depositStorage,
	}
}

func (d *DepositService) SaveDeposit(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint) error {

	return d.depositStorage.SaveDeposit(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount)
}

func (d *DepositService) GetDeposits(userId uint) ([]*handler.GetDepositResponse, error) {
	return d.getDepositsImpl(userId)
}

func (d *DepositService) getDepositsImpl(userId uint) ([]*handler.GetDepositResponse, error) {
	serviceDeposits, err := d.depositStorage.GetDeposits(userId)
	if err != nil {
		return nil, errors.New("cannot get deposits")
	}

	if len(serviceDeposits) == 0 {
		return nil, errors.New("there are no deposits")
	}

	handlerDeposits := d.serviceDepositsToGetDepositRes(serviceDeposits)

	return handlerDeposits, nil
}
