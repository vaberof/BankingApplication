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
