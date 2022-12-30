package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/service/transfer"
)

func RenderTransferResponse(c *fiber.Ctx, transfer *transfer.Transfer) error {
	return RenderResponse(c, transfer)
}

func RenderTransfersResponse(c *fiber.Ctx, transfers []*transfer.Transfer) error {
	return RenderResponse(c, transfers)
}
