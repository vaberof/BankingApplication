package deposit

import (
	"github.com/vaberof/MockBankingApplication/internal/app/http/handler"
)

func (d *DepositService) serviceDepositToGetDepositRes(serviceDeposit *Deposit) *handler.GetDepositResponse {
	var handlerDeposit handler.GetDepositResponse

	handlerDeposit.SenderUsername = serviceDeposit.SenderUsername
	handlerDeposit.SenderAccountId = serviceDeposit.SenderAccountId
	handlerDeposit.PayeeAccountId = serviceDeposit.PayeeAccountId
	handlerDeposit.Amount = serviceDeposit.Amount
	handlerDeposit.Date = serviceDeposit.Date

	return &handlerDeposit
}

func (d *DepositService) serviceDepositsToGetDepositRes(serviceDeposits []*Deposit) []*handler.GetDepositResponse {
	var handlerDeposits []*handler.GetDepositResponse

	for i := 0; i < len(serviceDeposits); i++ {
		serviceDeposit := serviceDeposits[i]
		handlerDeposits = append(handlerDeposits, d.serviceDepositToGetDepositRes(serviceDeposit))
	}

	return handlerDeposits
}
