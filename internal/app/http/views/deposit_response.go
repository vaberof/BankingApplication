package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/service/deposit"
	"time"
)

type DepositResponse struct {
	SenderUsername  string    `json:"sender_username"`
	SenderAccountId uint      `json:"sender_account_id"`
	PayeeAccountId  uint      `json:"payee_account_id"`
	Amount          uint      `json:"amount"`
	Date            time.Time `json:"date"`
}

func buildDeposit(deposit *deposit.Deposit) *DepositResponse {
	return &DepositResponse{
		SenderUsername:  deposit.SenderUsername,
		SenderAccountId: deposit.SenderAccountId,
		PayeeAccountId:  deposit.PayeeAccountId,
		Amount:          deposit.Amount,
		Date:            deposit.Date,
	}
}

func RenderDepositsResponse(c *fiber.Ctx, deposits []*deposit.Deposit) error {
	depositsResponse := make([]*DepositResponse, len(deposits))

	for i := 0; i < len(deposits); i++ {
		depositsResponse[i] = buildDeposit(deposits[i])
	}

	return RenderResponse(c, depositsResponse)
}
