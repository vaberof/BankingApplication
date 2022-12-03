package handler

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *HttpHandler) register(c *fiber.Ctx) error {
	var createUserReqBody CreateUserRequestBody

	err := c.BodyParser(&createUserReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	_, err = h.userService.GetUserByUsername(createUserReqBody.Username)
	if err == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "user with this username already exist",
		})
	}

	err = h.userService.CreateUser(createUserReqBody.Username, createUserReqBody.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": "cannot create user",
		})
	}

	user, err := h.userService.GetUserByUsername(createUserReqBody.Username)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = h.accountService.CreateInitialAccount(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "successfully created a user",
	})
}
