package handler

import "github.com/vaberof/MockBankingApplication/internal/service/transfer"

type TransferService interface {
	MakeTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*transfer.Transfer, error)
	GetTransfers(userId uint) ([]*transfer.Transfer, error)
}
