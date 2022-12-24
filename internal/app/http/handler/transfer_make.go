package handler

import "github.com/gofiber/fiber/v2"

type makeTransferRequestBody struct {
	SenderAccountId uint `json:"sender_account_id"`
	PayeeId         uint `json:"payee_id"`
	PayeeAccountId  uint `json:"payee_account_id"`
	Amount          uint `json:"amount"`
}

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
		makeTransferReqBody.PayeeId,
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
