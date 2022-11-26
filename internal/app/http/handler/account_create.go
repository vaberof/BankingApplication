package handler

import (
	"github.com/gofiber/fiber/v2"
)

type CreateAccountRequestBody struct {
	Name string `json:"name"`
}

func (h *Handler) createAccount(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var createAccountReqBody CreateAccountRequestBody

	err = c.BodyParser(&createAccountReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.accountService.CreateCustomAccount(user.Id, createAccountReqBody.Name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": "cannot create account",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"error": "successfully created an account",
	})
}
