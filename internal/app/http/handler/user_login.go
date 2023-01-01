package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type userLoginRequestBody struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

//	@Summary		Login
//	@Tags			Auth
//	@Description	Login into account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		userLoginRequestBody	true	"User data"
//	@Success		200		{string}	string					"Successfully logged in"
//	@Failure		400		{string}	error					"Invalid Request Body"
//	@Failure		401		{string}	error					"Authorization information is missing or invalid"
//	@Failure		500		{string}	error					"Unexpected error"
//	@Router			/auth/login [post]
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
