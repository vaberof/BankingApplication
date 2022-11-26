package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserLoginRequestBody struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

func (h *Handler) login(c *fiber.Ctx) error {
	var userLoginReqBody UserLoginRequestBody

	err := c.BodyParser(&userLoginReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.authService.GenerateJwtToken(userLoginReqBody.Username, userLoginReqBody.Password)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := h.authService.GenerateCookie(token)
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
