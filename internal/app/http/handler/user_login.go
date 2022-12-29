package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.authService.GenerateJwtToken(userLoginReqBody.Username, userLoginReqBody.Password)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := h.authService.GenerateCookie(token)
	c.Cookie(cookie)

	log.Printf("TOKEN: %s, COOKIE: %v, error in USER LOGIN: %v", token, c.Cookies("jwt"), err)

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"token": token,
	})
}
