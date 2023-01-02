package depositpg

import (
	service "github.com/vaberof/MockBankingApplication/internal/service/deposit"
)

func BuildServiceDeposit(postgresDeposit *PostgresDeposit) *service.Deposit {
	return &service.Deposit{
		SenderId:        postgresDeposit.SenderId,
		SenderUsername:  postgresDeposit.SenderUsername,
		SenderAccountId: postgresDeposit.SenderAccountId,
		PayeeId:         postgresDeposit.PayeeId,
		PayeeAccountId:  postgresDeposit.PayeeAccountId,
		Amount:          postgresDeposit.Amount,
		Date:            postgresDeposit.CreatedAt,
	}
}

func BuildServiceDeposits(postgresDeposits []*PostgresDeposit) []*service.Deposit {
	serviceTransfers := make([]*service.Deposit, len(postgresDeposits))

	for i := 0; i < len(serviceTransfers); i++ {
		serviceTransfers[i] = BuildServiceDeposit(postgresDeposits[i])
	}

	return serviceTransfers
}
