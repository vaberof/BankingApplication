package deposit

type DepositService struct {
	userStorage    UserStorage
	depositStorage DepositStorage
}

func NewDepositService(depositStorage DepositStorage, userStorage UserStorage) *DepositService {
	return &DepositService{
		depositStorage: depositStorage,
		userStorage:    userStorage,
	}
}

func (d *DepositService) SaveDeposit(
	senderId uint,
	senderAccountId uint,
	payeeId uint,
	payeeAccountId uint,
	amount uint,
	depositType string) error {

	return d.saveDepositImpl(senderId, senderAccountId, payeeId, payeeAccountId, amount, depositType)
}

func (d *DepositService) saveDepositImpl(
	senderId uint,
	senderAccountId uint,
	payeeId uint,
	payeeAccountId uint,
	amount uint,
	depositType string) error {

	user, err := d.userStorage.GetUserById(senderId)

	if err != nil {
		return err
	}

	return d.depositStorage.SaveDeposit(senderId, user.Username, senderAccountId, payeeId, payeeAccountId, amount, depositType)
}
