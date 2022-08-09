package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"github.com/vaberof/banking_app/internal/pkg/typeconv"
)

func (h *Handler) getTransfers(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")

	token, err := h.services.Authorization.ParseJwtToken(jwtToken)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	userId, err := typeconv.ConvertStringIdToUintId(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	transfers, err := h.services.Transfer.GetTransfers(userId)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.TransfersNotFound,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"transfers": transfers,
	})
}
