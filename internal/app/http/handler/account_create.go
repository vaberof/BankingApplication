package handler

import (
	"github.com/gofiber/fiber/v2"
)

type createAccountRequestBody struct {
	Name string `json:"name"`
}

//	@Summary		Create a bank account
//	@Tags			Bank Account
//	@Description	Create a new custom bank account with specific name
//	@ID				creates custom bank account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		createAccountRequestBody	true	"account name"
//	@Success		200		{string}	error
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Router			/account [post]
func (h *HttpHandler) createAccount(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var createAccountReqBody createAccountRequestBody

	err = c.BodyParser(&createAccountReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.accountService.CreateCustomAccount(user.Id, createAccountReqBody.Name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "successfully created an account",
	})
}
