package handler

import "github.com/gofiber/fiber/v2"

type deleteAccountRequestBody struct {
	Name string `json:"name"`
}

func (h *HttpHandler) deleteAccount(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	var deleteAccountReqBody deleteAccountRequestBody

	err = c.BodyParser(&deleteAccountReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err = h.accountService.DeleteAccount(user.Id, deleteAccountReqBody.Name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "account successfully deleted",
	})
}
