package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"github.com/vaberof/MockBankingApplication/internal/pkg/responses"
)

// @Summary SignUp
// @Tags Auth
// @Description Create a new user
// @ID creates new user
// @Accept json
// @Produce json
// @Param input body domain.User true "user info"
// @Success 200 {string} string responses.Success
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /signup [post]
func (h *Handler) signUp(c *fiber.Ctx) error {
	var input domain.User

	err := c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	if h.services.AuthorizationValidator.IsEmptyUsername(input.Username) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.EmptyUsername,
		})
	}

	if h.services.AuthorizationValidator.IsEmptyPassword(input.Password) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.EmptyPassword,
		})
	}

	if err = h.services.AuthorizationValidator.UserExists(input.Username); err == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.UserAlreadyExists,
		})
	}

	err = h.services.Authorization.CreateUser(&input)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.services.Account.CreateInitialAccount(input.Id, input.Username)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
