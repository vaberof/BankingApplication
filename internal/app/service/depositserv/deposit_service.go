package depositserv

import (
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"github.com/vaberof/MockBankingApplication/internal/app/repository"
)

type DepositService struct {
	rDeposit       repository.Deposit
	rUserFinder    repository.UserFinder
	rAccountFinder repository.AccountFinder
}

func NewDepositService(rDeposit repository.Deposit, rUserFinder repository.UserFinder, rAccountFinder repository.AccountFinder) *DepositService {
	return &DepositService{
		rDeposit:       rDeposit,
		rUserFinder:    rUserFinder,
		rAccountFinder: rAccountFinder,
	}
}

func (s *DepositService) ConvertTransferToDeposit(transfer *domain.Transfer) (*domain.Deposit, error) {
	deposit := domain.NewDeposit()

	senderUser, err := s.rUserFinder.GetUserById(transfer.SenderId)
	if err != nil {
		return nil, err
	}

	payeeAccount, err := s.rAccountFinder.GetAccountById(transfer.PayeeAccountId)
	if err != nil {
		return nil, err
	}

	payeeUser, err := s.rUserFinder.GetUserById(payeeAccount.UserId)
	if err != nil {
		return nil, err
	}

	deposit.SetSenderId(senderUser.Id)
	deposit.SetSenderUsername(senderUser.Username)
	deposit.SetSenderAccountId(transfer.SenderAccountId)
	deposit.SetPayeeId(payeeUser.Id)
	deposit.SetPayeeAccountId(transfer.PayeeAccountId)
	deposit.SetAmount(transfer.Amount)
	deposit.SetType(transfer.Type)

	return deposit, nil
}

func (s *DepositService) CreateDeposit(deposit *domain.Deposit) error {
	return s.rDeposit.CreateDeposit(deposit)
}

func (s *DepositService) GetDeposits(userId uint) (*domain.Deposits, error) {
	deposits, err := s.rDeposit.GetDeposits(userId)
	if err != nil {
		return nil, err
	}
	return deposits, nil
}
