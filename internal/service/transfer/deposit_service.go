package transfer

type DepositService interface {
	SaveDeposit(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint, depositType string) error
}
