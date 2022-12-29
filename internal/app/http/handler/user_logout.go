package handler

import (
	"github.com/gofiber/fiber/v2"
)

//	@Summary		Logout
//	@Tags			Auth
//	@Description	Logout from account
//	@ID				logs out from account
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		401	{string}	string
//	@Router			/logout [post]
func (h *HttpHandler) logout(c *fiber.Ctx) error {
	_, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	cookie := h.authService.RemoveCookie()
	c.Cookie(cookie)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "successfully logout",
	})
}
