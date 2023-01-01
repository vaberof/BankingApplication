package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type createAccountRequestBody struct {
	Name string `json:"name"`
}

//	@Summary		Create a bank account
//	@Tags			Bank Account
//	@Description	Create a new bank account with specific name
//	@Accept			json
//	@Produce		json
//	@Param			input	body		createAccountRequestBody	true	"Account name"
//	@Success		200		{object}	views.AccountResponse		"Successfully created"
//	@Failure		400		{string}	string						"Invalid request body"
//	@Failure		401		{string}	string						"Authorization information is missing or invalid"
//	@Failure		500		{string}	string						"Unexpected error"
//	@Router			/account [post]
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
