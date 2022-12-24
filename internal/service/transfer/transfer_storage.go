package transfer

type TransferStorage interface {
	MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint, transferType string) error
}
