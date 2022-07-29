package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/constants"
	"github.com/vaberof/banking_app/internal/app/service"
)

func CreateAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": constants.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var data map[string]string

	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	accountType := data["type"]

	if service.IsEmptyAccountType(accountType) {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": constants.EmptyAccountType,
		})
	}

	_, err = service.FindAccountByType(accountType, claims)
	if err == nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": constants.AccountAlreadyExists,
		})
	}

	account := service.CreateCustomAccount(accountType, claims)
	service.CreateAccountInDatabase(account)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": constants.Success,
	})
}

func DeleteAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": constants.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var data map[string]string

	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	accountType := data["type"]

	if service.IsMainAccountType(accountType) {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": constants.FailedDeleteMainAccount,
		})
	}

	account, err := service.FindAccountByType(accountType, claims)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": constants.AccountNotFound,
		})
	}

	if !service.IsZeroBalance(account) {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": constants.FailedDeleteNonZeroBalanceAccount,
		})
	}

	service.DeleteAccountFromDatabase(account)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": constants.Success,
	})
}
