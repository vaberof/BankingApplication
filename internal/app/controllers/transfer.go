package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"os"
)

func MakeTransfer(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}
	claims := token.Claims.(*jwt.RegisteredClaims)

	var data map[string]string

	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	err = service.MakeTransfer(data, claims)
	if err != nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.FailedTransfer,
			"error":   err.Error(),
		})
	}

	senderUserID, senderAccountID, payeeUsername, payeeAccountID, amount, transferType := service.GetTransferData(data, claims)
	transfer := service.CreateTransfer(senderUserID, senderAccountID, payeeUsername, payeeAccountID, amount, transferType)
	service.CreateTransferInDatabase(transfer)

	senderUsername, payeeID := service.GetDepositData(data, claims)
	deposit := service.CreateDeposit(senderUserID, senderUsername, senderAccountID, payeeID, payeeAccountID, amount, transferType)
	service.CreateDepositInDatabase(deposit)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
