package deposit

type DepositService struct {
	depositStorage DepositStorage
}

func NewDepositService(depositStorage DepositStorage) *DepositService {
	return &DepositService{
		depositStorage: depositStorage,
	}
}

func (d *DepositService) SaveDeposit(
	senderId uint,
	senderUsername string,
	senderAccountId uint,
	payeeId uint,
	payeeUsername string,
	payeeAccountId uint,
	amount uint,
	depositType string) error {

	return d.depositStorage.SaveDeposit(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount, depositType)
}
