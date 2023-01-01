package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

//	@Summary		Get all transfers
//	@Tags			Transfer
//	@Description	Get all transfers user have made between own/other accounts
//	@Produce		json
//	@Success		200	{array}		transfer.Transfer	"Successfully retrieved"
//
//	@Failure		401	{string}	error				"Authorization information is missing or invalid"
//	@Failure		500	{string}	error				"Unexpected error"
//
//	@Router			/transfers [get]
func (h *HttpHandler) getTransfers(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	transfers, err := h.transferService.GetTransfers(user.Id)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderTransfersResponse(c, transfers)
}
