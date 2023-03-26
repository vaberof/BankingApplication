package transferpg

import (
	service "github.com/vaberof/MockBankingApplication/internal/service/transfer"
)

func BuildServiceTransfer(postgresTransfer *PostgresTransfer) *service.Transfer {
	return &service.Transfer{
		SenderAccountId: postgresTransfer.SenderAccountId,
		PayeeAccountId:  postgresTransfer.PayeeAccountId,
		PayeeUsername:   postgresTransfer.PayeeUsername,
		Amount:          postgresTransfer.Amount,
		TransferType:    postgresTransfer.TransferType,
		Date:            postgresTransfer.CreatedAt,
	}
}

func BuildServiceTransfers(postgresTransfers []*PostgresTransfer) []*service.Transfer {
	serviceTransfers := make([]*service.Transfer, len(postgresTransfers))

	for i := 0; i < len(serviceTransfers); i++ {
		serviceTransfers[i] = BuildServiceTransfer(postgresTransfers[i])
	}

	return serviceTransfers
}
