package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/app/http/views"
)

type createUserRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary		Register
// @Tags			Auth
// @Description	Register new user
// @ID				Registers new user
// @Accept			json
// @Produce		json
// @Param			input	body		createUserRequestBody	true	"user data"
// @Success		200		{string}	error
// @Failure		400		{object}	error
// @Failure		500		{object}	error
// @Router			/register [post]
func (h *HttpHandler) register(c *fiber.Ctx) error {
	var createUserReqBody createUserRequestBody

	err := c.BodyParser(&createUserReqBody)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.userService.CreateUser(createUserReqBody.Username, createUserReqBody.Password)
	if err != nil {
		return views.RenderErrorResponse(c, fiber.StatusInternalServerError, err.Error())

	}

	return views.RenderUserResponse(c, user)
}
