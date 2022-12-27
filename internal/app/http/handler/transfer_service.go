package handler

import "github.com/vaberof/MockBankingApplication/internal/service/transfer"

type TransferService interface {
	MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error
	GetTransfers(userId uint) ([]*transfer.Transfer, error)
}
