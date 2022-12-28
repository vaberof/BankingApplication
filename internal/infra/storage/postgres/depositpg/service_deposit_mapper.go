package depositpg

import (
	"github.com/vaberof/MockBankingApplication/internal/service/deposit"
)

func (s *PostgresDepositStorage) infraDepositToService(infraDeposit *Deposit) *deposit.Deposit {
	var serviceDeposit deposit.Deposit

	serviceDeposit.SenderId = infraDeposit.SenderId
	serviceDeposit.SenderUsername = infraDeposit.SenderUsername
	serviceDeposit.SenderAccountId = infraDeposit.SenderAccountId
	serviceDeposit.PayeeId = infraDeposit.PayeeId
	serviceDeposit.PayeeAccountId = infraDeposit.PayeeAccountId
	serviceDeposit.Amount = infraDeposit.Amount
	serviceDeposit.Date = infraDeposit.CreatedAt

	return &serviceDeposit
}

func (s *PostgresDepositStorage) infraDepositsToService(infraDeposits []*Deposit) []*deposit.Deposit {
	var serviceDeposit []*deposit.Deposit

	for i := 0; i < len(infraDeposits); i++ {
		infraDeposit := infraDeposits[i]
		serviceDeposit = append(serviceDeposit, s.infraDepositToService(infraDeposit))
	}

	return serviceDeposit
}
