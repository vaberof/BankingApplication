package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func (h *Handler) Signup(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	inputUsername := data["username"]

	_, err = service.GetUser(data)
	if err == nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.UserAlreadyExists,
		})
	}

	inputPassword := data["password"]
	hashedPassword := service.HashPassword(inputPassword)

	user := service.CreateUser(inputUsername, hashedPassword)
	service.CreateUserInDatabase(user)

	userInitialAccount := service.CreateInitialAccount(user.ID, user.Username)
	service.CreateAccountInDatabase(userInitialAccount)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
