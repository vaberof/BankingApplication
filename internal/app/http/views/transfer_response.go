package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/service/transfer"
	"time"
)

type TransferResponse struct {
	SenderAccountId uint      `json:"sender_account_id"`
	PayeeAccountId  uint      `json:"payee_account_id"`
	PayeeUsername   string    `json:"payee_username"`
	Amount          uint      `json:"amount"`
	TransferType    string    `json:"transfer_type"`
	Date            time.Time `json:"date"`
}

func buildTransfer(transfer *transfer.Transfer) *TransferResponse {
	return &TransferResponse{
		SenderAccountId: transfer.SenderAccountId,
		PayeeAccountId:  transfer.PayeeAccountId,
		PayeeUsername:   transfer.PayeeUsername,
		Amount:          transfer.Amount,
		TransferType:    transfer.TransferType,
		Date:            transfer.Date,
	}
}

func RenderTransferResponse(c *fiber.Ctx, transfer *transfer.Transfer) error {
	return RenderResponse(c, buildTransfer(transfer))
}

func RenderTransfersResponse(c *fiber.Ctx, transfers []*transfer.Transfer) error {
	transfersResponse := make([]*TransferResponse, len(transfers))

	for i := 0; i < len(transfersResponse); i++ {
		transfersResponse[i] = buildTransfer(transfers[i])
	}

	return RenderResponse(c, transfers)
}
