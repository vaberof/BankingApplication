package transfer

import "errors"

type TransferService struct {
	transferStorage     TransferStorage
	depositService      DepositService
	accountStorage      AccountStorage
	userResponseService UserService
}

func NewTransferService(
	transferStorage TransferStorage,
	depositService DepositService,
	accountStorage AccountStorage,
	userResponseService UserService) *TransferService {

	return &TransferService{
		transferStorage:     transferStorage,
		depositService:      depositService,
		accountStorage:      accountStorage,
		userResponseService: userResponseService,
	}
}

func (s *TransferService) MakeTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	return s.makeTransferImpl(senderId, senderAccountId, payeeAccountId, amount)
}

func (s *TransferService) GetTransfers(userId uint) ([]*Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *TransferService) makeTransferImpl(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	if senderAccountId == payeeAccountId {
		return nil, errors.New("cannot make a transfer to the same account")
	}

	payeeId, senderUsername, payeeUsername, err := s.preprocessTransfer(senderId, payeeAccountId)
	if err != nil {
		return nil, err
	}

	transferType, err := s.processTransfer(senderId, payeeId, senderAccountId, payeeAccountId, amount)
	if err != nil {
		return nil, err
	}

	transfer, err := s.transferStorage.SaveTransfer(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount, transferType)
	if err != nil {
		return nil, err
	}

	if senderId == payeeId {
		return nil, nil
	}

	err = s.depositService.SaveDeposit(senderId, senderUsername, senderAccountId, payeeId, payeeUsername, payeeAccountId, amount)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}

func (s *TransferService) preprocessTransfer(senderId uint, payeeAccountId uint) (uint, string, string, error) {
	payeeAccount, err := s.accountStorage.GetAccountById(payeeAccountId)
	if err != nil {
		return 0, "", "", errors.New("cannot get payee`s account")
	}

	payeeUser, err := s.userResponseService.GetUserById(payeeAccount.UserId)
	if err != nil {
		return 0, "", "", errors.New("cannot get payee user")
	}

	if senderId == payeeUser.Id {
		return payeeUser.Id, payeeUser.Username, payeeUser.Username, nil
	}

	senderUser, err := s.userResponseService.GetUserById(senderId)
	if err != nil {
		return 0, "", "", errors.New("cannot get sender user")
	}

	return payeeUser.Id, senderUser.Username, payeeUser.Username, nil
}

func (s *TransferService) processTransfer(senderId uint, payeeId uint, senderAccountId uint, payeeAccountId uint, amount uint) (string, error) {
	if !s.isAcceptableAmount(amount) {
		return "", errors.New("amount must be greater than 0")
	}

	transferType := s.getTransferType(senderId, payeeId)

	return transferType, nil
}

func (s *TransferService) getTransfersImpl(userId uint) ([]*Transfer, error) {
	transfers, err := s.transferStorage.GetTransfers(userId)
	if err != nil {
		return nil, errors.New("cannot get transfers")
	}

	if len(transfers) == 0 {
		return nil, errors.New("you have not made any transfers yet")
	}

	return transfers, nil
}

func (s *TransferService) getTransferType(senderId uint, payeeId uint) string {
	if senderId == payeeId {
		return "personal"
	}
	return "client"
}

func (s *TransferService) isAcceptableAmount(amount uint) bool {
	return amount > 0
}
