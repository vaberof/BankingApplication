package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

// @Summary SignIn
// @Tags Auth
// @Description authorization
// @ID auth-user
// @Accept json
// @Produce json
// @Param input body domain.User true "user info"
// @Success 200 {string} string responses.Success
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /login [post]
func (h *Handler) login(c *fiber.Ctx) error {
	var input domain.User

	err := c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	user, err := h.services.UserFinder.GetUserByUsername(input.Username)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.UserNotfound,
		})
	}

	if err = h.services.AuthorizationValidator.IsCorrectPassword(user.Password, input.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.IncorrectUsernameAndOrPassword,
		})
	}

	token, err := h.services.Authorization.GenerateJwtToken(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedToGenerateJwtToken,
		})
	}

	cookie := h.services.Authorization.GenerateCookie(token)
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
