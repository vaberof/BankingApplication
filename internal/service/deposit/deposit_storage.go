package deposit

type DepositStorage interface {
	SaveDeposit(senderId uint, senderUsername string, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint, depositType string) error
}
