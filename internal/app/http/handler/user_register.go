package handler

import (
	"github.com/gofiber/fiber/v2"
)

type createUserRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *HttpHandler) register(c *fiber.Ctx) error {
	var createUserReqBody createUserRequestBody

	err := c.BodyParser(&createUserReqBody)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	userId, err := h.userService.CreateUser(createUserReqBody.Username, createUserReqBody.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "successfully created a user",
		"id":      userId,
	})
}
