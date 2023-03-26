package transfer

import "github.com/vaberof/MockBankingApplication/internal/domain/account"

type TransferStorage interface {
	SaveTransfer(
		senderUsername string,
		senderAccount *account.Account,
		payeeUsername string,
		payeeAccount *account.Account,
		amount uint,
		transferType string) (*Transfer, error)

	GetTransfers(userId uint) ([]*Transfer, error)
}
