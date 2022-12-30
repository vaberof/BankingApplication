package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

// @Summary		Get all bank accounts
// @Tags			Bank Account
// @Description	Get all bank accounts you have
// @ID				gets all bank accounts
// @Produce		json
// @Success		200	{array}		views.AccountResponse
// @Failure		401	{object}	error
// @Failure		500	{object}	error
// @Router			/accounts [get]
func (h *HttpHandler) getAccounts(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	accounts, err := h.accountService.GetAccounts(user.Id)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderAccountsResponse(c, accounts)
}
