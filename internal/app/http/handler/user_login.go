package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type userLoginRequestBody struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

// @Summary		Login
// @Tags			Auth
// @Description	Login into account
// @ID				logins into account
// @Accept			json
// @Produce		json
// @Param			input	body		userLoginRequestBody	true	"user data"
// @Success		200		{string}	error
// @Failure		400		{object}	error
// @Failure		401		{object}	error
// @Failure		500		{object}	error
// @Router			/login [post]
func (h *HttpHandler) login(c *fiber.Ctx) error {
	var userLoginReqBody userLoginRequestBody

	err := c.BodyParser(&userLoginReqBody)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	token, err := h.authService.GenerateJwtToken(userLoginReqBody.Username, userLoginReqBody.Password)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	cookie := h.authService.GenerateCookie(token)
	c.Cookie(cookie)

	return views.RenderResponse(c, token)
}
