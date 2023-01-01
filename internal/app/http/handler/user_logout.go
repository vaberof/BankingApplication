package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

//	@Summary		Logout
//	@Tags			Auth
//	@Description	Logout from account
//	@Produce		json
//	@Success		200	{string}	string	"Successfully logged out"
//
//	@Failure		401	{string}	error	"Authorization information is missing or invalid"
//
//	@Router			/auth/logout [post]
func (h *HttpHandler) logout(c *fiber.Ctx) error {
	_, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	cookie := h.authService.RemoveCookie()
	c.Cookie(cookie)

	return views.RenderResponse(c, views.ResponseMessage{
		"message": "successfully logged out"})
}
