package transfer

type TransferStorage interface {
	SaveTransfer(
		senderId uint,
		senderUsername string,
		senderAccountId uint,
		payeeId uint,
		payeeUsername string,
		payeeAccountId uint,
		amount uint,
		transferType string) (*Transfer, error)

	GetTransfers(userId uint) ([]*Transfer, error)
}
