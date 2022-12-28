package deposit

type DepositStorage interface {
	SaveDeposit(
		senderId uint,
		senderUsername string,
		senderAccountId uint,
		payeeId uint,
		payeeUsername string,
		payeeAccountId uint,
		amount uint) error

	GetDeposits(userId uint) ([]*Deposit, error)
}
