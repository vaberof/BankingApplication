package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type deleteAccountRequestBody struct {
	Name string `json:"name"`
}

//	@Summary		Delete a bank account
//	@Tags			Bank Account
//	@Description	Delete a bank account with specific name
//	@Accept			json
//	@Produce		json
//	@Param			input	body		deleteAccountRequestBody	true	"Account name"
//	@Success		200		{string}	string						"Successfully deleted"
//	@Failure		400		{string}	error						"Invalid Request Body"
//
//	@Failure		401		{string}	error						"Authorization information is missing or invalid"
//	@Failure		500		{string}	error						"Unexpected error"
//
//	@Router			/account [delete]
func (h *HttpHandler) deleteAccount(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var deleteAccountReqBody deleteAccountRequestBody

	err = c.BodyParser(&deleteAccountReqBody)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	err = h.accountService.DeleteAccount(user.Id, deleteAccountReqBody.Name)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderResponse(c, views.ResponseMessage{
		"message": "account successfully deleted"})
}
