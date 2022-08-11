package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/pkg/responses"
)

// @Summary Logout
// @Tags Auth
// @Description exits from acc
// @ID logs out from account
// @Produce json
// @Success 200 {string} string responses.Success
// @Failure 401 {string} string responses.Unauthorized
// @Router /logout [post]
func (h *Handler) logout(c *fiber.Ctx) error {
	cookie := h.services.Authorization.RemoveCookie()
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
