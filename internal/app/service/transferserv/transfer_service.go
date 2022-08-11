package transferserv

import (
	"errors"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/app/repository"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

type TransferService struct {
	rTransfer        repository.Transfer
	rTransferAccount repository.TransferAccount
	rAccountFinder   repository.AccountFinder
}

func NewTransferService(rTransfer repository.Transfer, rTransferAccount repository.TransferAccount, rAccountFinder repository.AccountFinder) *TransferService {
	return &TransferService{
		rTransfer:        rTransfer,
		rTransferAccount: rTransferAccount,
		rAccountFinder:   rAccountFinder,
	}
}

func (s *TransferService) TransformInputToTransfer(
	senderId uint,
	senderAccountId uint,
	payeeAccountId uint,
	amount int,
	transferType string) (*domain.Transfer, error) {

	transfer := domain.NewTransfer()

	payeeAccount, err := s.rAccountFinder.GetAccountById(payeeAccountId)
	if err != nil {
		customError := errors.New(responses.PayeeAccountNotFound)
		return nil, customError
	}

	transfer.SetSenderId(senderId)
	transfer.SetSenderAccountId(senderAccountId)
	transfer.SetPayeeUsername(payeeAccount.Owner)
	transfer.SetPayeeAccountId(payeeAccountId)
	transfer.SetAmount(amount)
	transfer.SetType(transferType)

	return transfer, nil
}

func (s *TransferService) MakeTransfer(transfer *domain.Transfer) error {
	transferType := transfer.Type
	switch transferType {
	case "client":
		return s.clientTransfer(transfer)
	case "personal":
		return s.personalTransfer(transfer)
	default:
		customError := errors.New(responses.UnsupportedTransferType)
		return customError
	}
}

func (s *TransferService) CreateTransfer(transfer *domain.Transfer) error {
	return s.rTransfer.CreateTransfer(transfer)
}

func (s *TransferService) GetTransfers(userId uint) (*domain.Transfers, error) {
	transfers, err := s.rTransfer.GetTransfers(userId)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func (s *TransferService) clientTransfer(transfer *domain.Transfer) error {
	senderAccountId := transfer.SenderAccountId
	payeeAccountId := transfer.PayeeAccountId
	amount := transfer.Amount

	if isZeroAmountTransfer(amount) {
		customError := errors.New(responses.CannotTransferZeroAmount)
		return customError
	}

	senderAccount, err := s.rTransferAccount.GetSenderAccount(transfer.SenderId, senderAccountId)
	if err != nil {
		return err
	}

	payeeAccount, err := s.rTransferAccount.GetClientPayeeAccount(payeeAccountId)
	if err != nil {
		return err
	}

	if isSameAccountOwner(senderAccount.UserId, payeeAccount.UserId) {
		customError := errors.New(responses.SenderIsPayee)
		return customError
	}

	senderBalance := senderAccount.Balance
	payeeBalance := payeeAccount.Balance

	if !isEnoughFunds(senderBalance, amount) {
		customError := errors.New(responses.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderBalance - amount
	newPayeeBalance := payeeBalance + amount

	senderAccountDbObject, err := s.rTransferAccount.GetAccountDbObject(senderAccountId)
	if err != nil {
		return err
	}

	payeeAccountDbObject, err := s.rTransferAccount.GetAccountDbObject(payeeAccountId)
	if err != nil {
		return err
	}

	err = s.rTransferAccount.UpdateAccountBalanceDbObject(senderAccountDbObject, newSenderBalance)
	if err != nil {
		return err
	}

	err = s.rTransferAccount.UpdateAccountBalanceDbObject(payeeAccountDbObject, newPayeeBalance)
	if err != nil {
		return err
	}

	return nil
}

func (s *TransferService) personalTransfer(transfer *domain.Transfer) error {
	userId := transfer.SenderId

	senderAccountId := transfer.SenderAccountId
	payeeAccountId := transfer.PayeeAccountId
	amount := transfer.Amount

	if isZeroAmountTransfer(amount) {
		customError := errors.New(responses.CannotTransferZeroAmount)
		return customError
	}

	senderAccount, err := s.rTransferAccount.GetSenderAccount(userId, senderAccountId)
	if err != nil {
		return err
	}

	payeeAccount, err := s.rTransferAccount.GetPersonalPayeeAccount(userId, payeeAccountId)
	if err != nil {
		return err
	}

	if isSameAccountId(senderAccountId, payeeAccountId) {
		customError := errors.New(responses.SenderIsPayee)
		return customError
	}

	senderBalance := senderAccount.Balance
	payeeBalance := payeeAccount.Balance

	if !isEnoughFunds(senderBalance, amount) {
		customError := errors.New(responses.InsufficientFunds)
		return customError
	}

	newSenderBalance := senderBalance - amount
	newPayeeBalance := payeeBalance + amount

	senderAccountDbObject, err := s.rTransferAccount.GetAccountDbObject(senderAccountId)
	if err != nil {
		return err
	}

	payeeAccountDbObject, err := s.rTransferAccount.GetAccountDbObject(payeeAccountId)
	if err != nil {
		return err
	}

	err = s.rTransferAccount.UpdateAccountBalanceDbObject(senderAccountDbObject, newSenderBalance)
	if err != nil {
		return err
	}

	err = s.rTransferAccount.UpdateAccountBalanceDbObject(payeeAccountDbObject, newPayeeBalance)
	if err != nil {
		return err
	}

	return nil
}

func isSameAccountId(senderAccountId uint, payeeAccountId uint) bool {
	return senderAccountId == payeeAccountId
}

func isSameAccountOwner(senderUserId uint, payeeAccountId uint) bool {
	return senderUserId == payeeAccountId
}

func isZeroAmountTransfer(amount int) bool {
	return amount <= 0
}

func isEnoughFunds(balance int, amount int) bool {
	return balance-amount >= 0
}
