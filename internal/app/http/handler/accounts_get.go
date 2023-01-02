package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

//	@Summary		Get all bank accounts
//	@Tags			Bank Account
//	@Description	Get all bank accounts user have
//	@Produce		json
//	@Success		200	{array}		views.AccountResponse	"Successfully retrieved"
//
//	@Failure		401	{string}	string					"Authorization information is missing or invalid"
//	@Failure		500	{string}	string					"Unexpected error"
//
//	@Router			/accounts [get]
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
