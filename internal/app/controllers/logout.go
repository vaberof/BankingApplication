package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/service"
)

func Logout(c *fiber.Ctx) error {
	cookie := service.RemoveCookie()
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": constants.Success,
	})
}
