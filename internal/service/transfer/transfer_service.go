package transfer

import (
	"errors"
	domain "github.com/vaberof/MockBankingApplication/internal/domain/user"
)

type TransferService struct {
	transferStorage TransferStorage
	depositService  DepositService
	accountStorage  AccountStorage
	userService     UserService
}

func NewTransferService(
	transferStorage TransferStorage,
	depositService DepositService,
	accountStorage AccountStorage,
	userResponseService UserService) *TransferService {

	return &TransferService{
		transferStorage: transferStorage,
		depositService:  depositService,
		accountStorage:  accountStorage,
		userService:     userResponseService,
	}
}

func (s *TransferService) MakeTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	return s.makeTransferImpl(senderId, senderAccountId, payeeAccountId, amount)
}

func (s *TransferService) GetTransfers(userId uint) ([]*Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *TransferService) makeTransferImpl(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	if err := s.preprocessTransfer(senderAccountId, payeeAccountId, amount); err != nil {
		return nil, err
	}

	senderUser, payeeUser, transferType, err := s.processTransfer(senderId, payeeAccountId)
	if err != nil {
		return nil, err
	}

	transfer, err := s.transferStorage.SaveTransfer(
		senderId,
		senderUser.Username,
		senderAccountId,
		payeeUser.Id,
		payeeUser.Username,
		payeeAccountId,
		amount,
		transferType)
	if err != nil {
		return nil, err
	}

	if s.isPersonalTransfer(senderId, payeeUser.Id) {
		return transfer, nil
	}

	err = s.depositService.SaveDeposit(
		senderId,
		senderUser.Username,
		senderAccountId,
		payeeUser.Id,
		payeeUser.Username,
		payeeAccountId,
		amount)
	if err != nil {
		return transfer, err
	}

	return transfer, nil
}

func (s *TransferService) preprocessTransfer(senderAccountId uint, payeeAccountId uint, amount uint) error {
	if senderAccountId == payeeAccountId {
		return errors.New("cannot make a transfer to the same account")
	}

	if !s.isAcceptableAmount(amount) {
		return errors.New("amount must be greater than 0")
	}

	return nil
}

func (s *TransferService) processTransfer(senderId uint, payeeAccountId uint) (*domain.User, *domain.User, string, error) {
	payeeAccount, err := s.accountStorage.GetAccountById(payeeAccountId)
	if err != nil {
		return nil, nil, "", errors.New("cannot get payee`s account")
	}

	payeeUser, err := s.userService.GetUserById(payeeAccount.UserId)
	if err != nil {
		return nil, nil, "", errors.New("cannot get payee user")
	}

	transferType := s.getTransferType(senderId, payeeAccount.UserId)

	if s.isPersonalTransfer(senderId, payeeUser.Id) {
		return payeeUser, payeeUser, transferType, nil
	}

	senderUser, err := s.userService.GetUserById(senderId)
	if err != nil {
		return nil, nil, "", errors.New("cannot get sender user")
	}

	return senderUser, payeeUser, transferType, nil
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

func (s *TransferService) isPersonalTransfer(senderId uint, payeeId uint) bool {
	return senderId == payeeId
}

func (s *TransferService) isAcceptableAmount(amount uint) bool {
	return amount > 0
}
