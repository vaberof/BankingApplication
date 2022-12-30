package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type deleteAccountRequestBody struct {
	Name string `json:"name"`
}

// @Summary		Delete a bank account
// @Tags			Bank Account
// @Description	Delete a bank account with specific name
// @ID				deletes custom bank account
// @Accept			json
// @Produce		json
// @Param			input	body		deleteAccountRequestBody	true	"account name"
// @Success		200		{string}	error
// @Failure		400		{object}	error
// @Failure		401		{object}	error
// @Failure		500		{object}	error
// @Router			/account [delete]
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
