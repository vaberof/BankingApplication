package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	user, err := service.GetUser(data)
	if err != nil {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.IncorrectUsernameAndOrPassword,
		})
	}

	inputPassword := data["password"]

	if !service.IsCorrectPassword(user.Password, inputPassword) {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": responses.IncorrectUsernameAndOrPassword,
		})
	}

	claims := service.CreateJwtClaims(user)

	token, err := service.CreateJwtToken(claims)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedLogin,
		})
	}

	cookie := service.CreateCookie(token)
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
