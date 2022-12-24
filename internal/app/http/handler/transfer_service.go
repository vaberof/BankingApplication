package handler

type TransferService interface {
	MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error
}
