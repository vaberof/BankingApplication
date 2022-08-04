package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/service"
)

func GetDeposits(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": constants.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	deposits, err := service.GetUserDeposits(claims)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": constants.DepositsNotFound,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"deposits": deposits,
	})
}
