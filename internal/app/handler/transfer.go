package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"github.com/vaberof/banking_app/internal/pkg/typeconv"
)

type inputTransfer struct {
	SenderAccountId uint   `json:"sender_account_id"`
	PayeeAccountId  uint   `json:"payee_account_id"`
	Amount          int    `json:"amount"`
	Type            string `json:"transfer_type"`
}

func (h *Handler) transfer(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")

	token, err := h.services.Authorization.ParseJwtToken(jwtToken)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var input inputTransfer

	err = c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	senderId, err := typeconv.ConvertStringIdToUintId(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	transfer, err := h.services.Transfer.TransformInputToTransfer(
		senderId,
		input.SenderAccountId,
		input.PayeeAccountId,
		input.Amount,
		input.Type)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.services.Transfer.MakeTransfer(transfer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedTransfer,
			"error":   err.Error(),
		})
	}

	err = h.services.Transfer.CreateTransfer(transfer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	deposit, err := h.services.Deposit.ConvertTransferToDeposit(transfer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.services.Deposit.CreateDeposit(deposit)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
