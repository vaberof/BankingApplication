package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

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

	accountType := data["type"]

	if service.IsEmptyAccountType(accountType) {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.EmptyAccountType,
		})
	}

	_, err = service.FindAccountByType(accountType, claims)
	if err == nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.AccountAlreadyExists,
		})
	}

	account := service.CreateCustomAccount(accountType, claims)
	service.CreateAccountInDatabase(account)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}

func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := service.ParseJwtToken(cookie)

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

	accountType := data["type"]

	if service.IsMainAccountType(accountType) {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteMainAccount,
		})
	}

	account, err := service.FindAccountByType(accountType, claims)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.AccountNotFound,
		})
	}

	if !service.IsZeroBalance(account) {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteNonZeroBalanceAccount,
		})
	}

	service.DeleteAccountFromDatabase(account)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
