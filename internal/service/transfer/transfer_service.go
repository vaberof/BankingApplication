package transfer

import "errors"

type TransferService struct {
	transferStorage TransferStorage
}

func NewTransferService(transferStorage TransferStorage) *TransferService {
	return &TransferService{transferStorage: transferStorage}
}

func (s *TransferService) MakeTransfer(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error {
	return s.makeTransferImpl(senderId, senderAccountId, payeeId, payeeAccountId, amount)
}

func (s *TransferService) makeTransferImpl(senderId uint, senderAccountId uint, payeeId uint, payeeAccountId uint, amount uint) error {
	transferType, err := s.preprocessTransfer(senderId, payeeId, amount)
	if err != nil {
		return err
	}

	return s.transferStorage.MakeTransfer(senderId, senderAccountId, payeeId, payeeAccountId, amount, transferType)
}

func (s *TransferService) preprocessTransfer(senderId uint, payeeId uint, amount uint) (string, error) {
	if !s.isAcceptableAmount(amount) {
		return "", errors.New("amount should be greater than 0")
	}

	transferType := s.getTransferType(senderId, payeeId)

	return transferType, nil
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
