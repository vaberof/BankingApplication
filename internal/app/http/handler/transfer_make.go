package handler

import "github.com/gofiber/fiber/v2"

type makeTransferRequestBody struct {
	SenderAccountId uint `json:"sender_account_id"`
	PayeeAccountId  uint `json:"payee_account_id"`
	Amount          uint `json:"amount"`
}

//	@Summary		Make a transfer
//	@Tags			Transfer
//	@Description	Make a transfer between your/other clients accounts
//	@ID				Makes a transfer
//	@Accept			json
//	@Produce		json
//	@Param			input	body		makeTransferRequestBody	true	"transfer data"
//	@Success		200		{string}	error
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/transfer [post]
func (h *HttpHandler) makeTransfer(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var makeTransferReqBody makeTransferRequestBody

	err = c.BodyParser(&makeTransferReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.transferService.MakeTransfer(
		user.Id,
		makeTransferReqBody.SenderAccountId,
		makeTransferReqBody.PayeeAccountId,
		makeTransferReqBody.Amount)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "cannot make a transfer",
			"error":   err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"error": "successfully made a transfer",
	})
}
