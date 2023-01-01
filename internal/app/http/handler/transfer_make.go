package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type makeTransferRequestBody struct {
	SenderAccountId uint `json:"sender_account_id"`
	PayeeAccountId  uint `json:"payee_account_id"`
	Amount          uint `json:"amount"`
}

//	@Summary		Make a transfer
//	@Tags			Transfer
//	@Description	Make a transfer between own/clients accounts
//	@Accept			json
//	@Produce		json
//	@Param			input	body		makeTransferRequestBody	true	"Transfer data"
//	@Success		200		{object}	views.TransferResponse	"Successfully made a transfer"
//	@Failure		400		{string}	string					"Invalid request body"
//
//	@Failure		401		{string}	string					"Authorization information is missing or invalid"
//	@Failure		500		{string}	string					"Unexpected error"
//
//	@Router			/transfer [post]
func (h *HttpHandler) makeTransfer(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var makeTransferReqBody makeTransferRequestBody

	err = c.BodyParser(&makeTransferReqBody)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	transfer, err := h.transferService.MakeTransfer(
		user.Id,
		makeTransferReqBody.SenderAccountId,
		makeTransferReqBody.PayeeAccountId,
		makeTransferReqBody.Amount)

	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return views.RenderTransferResponse(c, transfer)
}
