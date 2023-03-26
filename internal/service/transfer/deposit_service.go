package transfer

type DepositService interface {
	SaveDeposit(
		senderId uint,
		senderUsername string,
		senderAccountId uint,
		payeeId uint,
		payeeUsername string,
		payeeAccountId uint,
		amount uint) error
}
