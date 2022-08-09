package service

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

func (s *TransferService) MakeTransfer(userId uint, transfer *domain.Transfer) error {
	transferType := transfer.Type
	switch transferType {
	case "client":
		return s.clientTransfer(userId, transfer)
	case "personal":
		return s.personalTransfer(userId, transfer)
	default:
		customError := errors.New(responses.UnsupportedTransferType)
		return customError
	}
}

func (s *TransferService) CreateTransfer(userId uint, transfer *domain.Transfer) error {
	trans := domain.NewTransfer()

	payeeAccount, err := s.rAccountFinder.GetAccountById(transfer.PayeeAccountId)
	if err != nil {
		customError := errors.New(responses.AccountNotFound)
		return customError
	}

	trans.SetSenderId(userId)
	trans.SetSenderAccountId(transfer.SenderAccountId)
	trans.SetPayeeUsername(payeeAccount.Owner)
	trans.SetPayeeAccountId(payeeAccount.Id)
	trans.SetAmount(transfer.Amount)
	trans.SetType(transfer.Type)

	return s.rTransfer.CreateTransfer(trans)
}

func (s *TransferService) GetTransfers(userId uint) (*domain.Transfers, error) {
	transfers, err := s.rTransfer.GetTransfers(userId)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func (s *TransferService) clientTransfer(userId uint, transfer *domain.Transfer) error {
	senderAccountId := transfer.SenderAccountId
	payeeAccountId := transfer.PayeeAccountId
	amount := transfer.Amount

	senderAccount, err := s.rTransferAccount.GetSenderAccount(userId, senderAccountId)
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

func (s *TransferService) personalTransfer(userId uint, transfer *domain.Transfer) error {
	senderAccountId := transfer.SenderAccountId
	payeeAccountId := transfer.PayeeAccountId
	amount := transfer.Amount

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

func isEnoughFunds(balance int, amount int) bool {
	return balance-amount >= 0
}
