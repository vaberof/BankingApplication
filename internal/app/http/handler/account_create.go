package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type createAccountRequestBody struct {
	Name string `json:"name"`
}

// @Summary		Create a bank account
// @Tags			Bank Account
// @Description	Create a new custom bank account with specific name
// @ID				creates custom bank account
// @Accept			json
// @Produce		json
// @Param			input	body		createAccountRequestBody	true	"account name"
// @Success		200		{object}	views.AccountResponse
// @Failure		400		{object}	error
// @Failure		401		{object}	error
// @Failure		500		{object}	error
// @Router			/account [post]
func (h *HttpHandler) createAccount(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var createAccountReqBody createAccountRequestBody

	err = c.BodyParser(&createAccountReqBody)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	account, err := h.accountService.CreateCustomAccount(user.Id, createAccountReqBody.Name)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderAccountResponse(c, account)
}
