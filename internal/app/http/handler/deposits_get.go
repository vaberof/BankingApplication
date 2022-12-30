package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

// @Summary		Get all deposits
// @Tags			Deposit
// @Description	Get all deposits other clients have made to your accounts
// @ID				Gets deposits
// @Produce		json
// @Success		200	{array}		views.DepositResponse
// @Failure		401	{object}	error
// @Failure		500	{object}	error
// @Router			/deposits [get]
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
