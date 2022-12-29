package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type GetDepositResponse struct {
	SenderUsername  string    `json:"sender_username"`
	SenderAccountId uint      `json:"sender_account_id"`
	PayeeAccountId  uint      `json:"payee_account_id"`
	Amount          uint      `json:"amount"`
	Date            time.Time `json:"date"`
}

//	@Summary		Get all deposits
//	@Tags			Deposit
//	@Description	Get all deposits other clients have made to your accounts
//	@ID				Gets deposits
//	@Produce		json
//	@Success		200	{string}	error
//	@Failure		401	{object}	error
//	@Failure		500	{object}	error
//	@Router			/deposits [get]
func (h *HttpHandler) getDeposits(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	deposits, err := h.depositService.GetDeposits(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"deposits": deposits,
	})
}
