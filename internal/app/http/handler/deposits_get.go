package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

//	@Summary		Get all deposits
//	@Tags			Deposit
//	@Description	Get all deposits that have been made to user accounts from other clients
//	@Produce		json
//	@Success		200	{array}		views.DepositResponse	"Successfully retrieved"
//
//	@Failure		401	{string}	error					"Authorization information is missing or invalid"
//	@Failure		500	{string}	error					"Unexpected error"
//
//	@Router			/deposits [get]
func (h *HttpHandler) getDeposits(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	deposits, err := h.depositService.GetDeposits(user.Id)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderDepositsResponse(c, deposits)
}
