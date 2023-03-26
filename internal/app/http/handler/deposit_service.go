package handler

import "github.com/vaberof/MockBankingApplication/internal/service/deposit"

type DepositService interface {
	GetDeposits(userId uint) ([]*deposit.Deposit, error)
}
