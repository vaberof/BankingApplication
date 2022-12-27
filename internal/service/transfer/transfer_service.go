package transfer

import "errors"

type TransferService struct {
	transferStorage TransferStorage
	depositService  DepositService
}

func NewTransferService(transferStorage TransferStorage, depositService DepositService) *TransferService {
	return &TransferService{
		transferStorage: transferStorage,
		depositService:  depositService,
	}
}

func (s *TransferService) MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error {
	return s.makeTransferImpl(senderId, senderAccountId, payeeId, payeeAccountId, amount)
}

func (s *TransferService) GetTransfers(userId uint) ([]*Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *TransferService) makeTransferImpl(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error {
	transferType, err := s.preprocessTransfer(senderId, payeeId, amount)
	if err != nil {
		return err
	}

	err = s.transferStorage.MakeTransfer(senderId, senderAccountId, payeeId, payeeAccountId, amount, transferType)
	if err != nil {
		return err
	}

	err = s.depositService.SaveDeposit(senderId, senderAccountId, payeeId, payeeAccountId, amount, transferType)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransferService) preprocessTransfer(senderId uint, payeeId uint, amount uint) (string, error) {
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
