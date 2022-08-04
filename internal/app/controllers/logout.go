package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func Logout(c *fiber.Ctx) error {
	cookie := service.RemoveCookie()
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
