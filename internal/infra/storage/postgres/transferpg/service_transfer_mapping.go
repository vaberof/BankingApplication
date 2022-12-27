package transferpg

import (
	"github.com/vaberof/MockBankingApplication/internal/service/transfer"
)

func (s *PostgresTransferStorage) infraTransferToService(infraTransfer *Transfer) *transfer.Transfer {
	var serviceTransfer transfer.Transfer

	serviceTransfer.SenderAccountId = infraTransfer.SenderAccountId
	serviceTransfer.PayeeId = infraTransfer.PayeeId
	serviceTransfer.PayeeAccountId = infraTransfer.PayeeAccountId
	serviceTransfer.Amount = infraTransfer.Amount
	serviceTransfer.TransferType = infraTransfer.TransferType

	return &serviceTransfer
}

func (s *PostgresTransferStorage) infraTransfersToService(infraTransfers []*Transfer) []*transfer.Transfer {
	var serviceTransfers []*transfer.Transfer

	for i := 0; i < len(infraTransfers); i++ {
		infraAccount := infraTransfers[i]
		serviceTransfers = append(serviceTransfers, s.infraTransferToService(infraAccount))
	}

	return serviceTransfers
}
