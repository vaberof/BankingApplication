package handler

type DepositService interface {
	GetDeposits(userId uint) ([]*GetDepositResponse, error)
}
