package transfer

import (
	"errors"
	"fmt"
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
	"github.com/vaberof/MockBankingApplication/internal/domain/user"
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
	userService UserService) *TransferService {

	return &TransferService{
		transferStorage: transferStorage,
		depositService:  depositService,
		accountStorage:  accountStorage,
		userService:     userService,
	}
}

func (s *TransferService) MakeTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	return s.makeTransferImpl(senderId, senderAccountId, payeeAccountId, amount)
}

func (s *TransferService) GetTransfers(userId uint) ([]*Transfer, error) {
	return s.getTransfersImpl(userId)
}

func (s *TransferService) makeTransferImpl(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*Transfer, error) {
	senderAccount, err := s.preprocessTransfer(senderId, senderAccountId, payeeAccountId, amount)
	if err != nil {
		return nil, err
	}

	senderUser, payeeUser, payeeAccount, transferType, err := s.processTransfer(senderId, payeeAccountId)
	if err != nil {
		return nil, err
	}

	transfer, err := s.transferStorage.SaveTransfer(
		senderUser.Username,
		senderAccount,
		payeeUser.Username,
		payeeAccount,
		amount,
		transferType)
	if err != nil {
		return nil, err
	}

	if s.isPersonalTransfer(transferType) {
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
		return nil, err
	}

	return transfer, nil
}

func (s *TransferService) preprocessTransfer(senderId uint, senderAccountId uint, payeeAccountId uint, amount uint) (*account.Account, error) {
	senderAccount, err := s.accountStorage.GetAccountById(senderAccountId)
	if err != nil {
		return nil, errors.New("sender`s account not found")
	}

	if !s.isSenderAccountOwner(senderId, senderAccount.UserId) {
		return nil, errors.New("account does not belong to sender")
	}

	if s.isTheSameAccount(senderAccountId, payeeAccountId) {
		return nil, errors.New("cannot make a transfer to the same account")
	}

	if !s.isAcceptableAmount(amount) {
		return nil, errors.New("amount must be greater than 0")
	}

	if !s.isEnoughFunds(senderAccount.Balance, amount) {
		return nil, errors.New("insufficient funds to make a transfer")
	}

	return senderAccount, nil
}

func (s *TransferService) processTransfer(senderId uint, payeeAccountId uint) (*user.User, *user.User, *account.Account, string, error) {
	payeeAccount, err := s.accountStorage.GetAccountById(payeeAccountId)
	if err != nil {
		return nil, nil, nil, "", errors.New("cannot find payee`s account")
	}

	payeeUser, err := s.userService.GetUserById(payeeAccount.UserId)
	if err != nil {
		return nil, nil, nil, "", errors.New("cannot find payee user")
	}

	transferType := s.getTransferType(senderId, payeeAccount.UserId)

	if s.isPersonalTransfer(transferType) {
		return payeeUser, payeeUser, payeeAccount, transferType, nil
	}

	senderUser, err := s.userService.GetUserById(senderId)
	if err != nil {
		return nil, nil, nil, "", errors.New("cannot find sender user")
	}

	return senderUser, payeeUser, payeeAccount, transferType, nil
}

func (s *TransferService) getTransfersImpl(userId uint) ([]*Transfer, error) {
	transfers, err := s.transferStorage.GetTransfers(userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get transfers: %s", err.Error())
	}

	if len(transfers) == 0 {
		return nil, errors.New("there are no transfers yet")
	}

	return transfers, nil
}

func (s *TransferService) getTransferType(senderId uint, payeeId uint) string {
	if senderId == payeeId {
		return "personal"
	}
	return "client"
}

func (s *TransferService) isPersonalTransfer(transferType string) bool {
	return transferType == "personal"
}

func (s *TransferService) isTheSameAccount(senderAccountId uint, payeeAccountId uint) bool {
	return senderAccountId == payeeAccountId
}

func (s *TransferService) isSenderAccountOwner(senderId uint, accountOwnerId uint) bool {
	return senderId == accountOwnerId
}

func (s *TransferService) isAcceptableAmount(amount uint) bool {
	return amount > 0
}

func (s *TransferService) isEnoughFunds(balance int, amount uint) bool {
	return balance-int(amount) >= 0
}
