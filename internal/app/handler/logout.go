package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func (h *Handler) logout(c *fiber.Ctx) error {
	cookie := h.services.Authorization.RemoveCookie()
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
