package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"github.com/vaberof/banking_app/internal/pkg/typeconv"
)

type inputAccount struct {
	Type string `json:"type"`
}

// @Summary Create a bank account
// @Tags Bank Account
// @Description Create a new bank account
// @ID creates bank account
// @Accept json
// @Produce json
// @Param input body inputAccount true "account type"
// @Success 200 {string} string responses.Success
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /account [post]
func (h *Handler) createAccount(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")

	token, err := h.services.Authorization.ParseJwtToken(jwtToken)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var input inputAccount

	err = c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	if h.services.AccountValidator.IsEmptyAccountType(input.Type) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.EmptyAccountType,
		})
	}

	userId, err := typeconv.ConvertStringIdToUintId(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err = h.services.AccountValidator.AccountExists(userId, input.Type); err == nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.AccountAlreadyExists,
		})
	}

	user, err := h.services.UserFinder.GetUserById(userId)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.UserNotfound,
		})
	}

	err = h.services.Account.CreateCustomAccount(user.Id, user.Username, input.Type)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedCreateAccount,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}

// @Summary Delete a bank account
// @Tags Bank Account
// @Description Delete a bank account
// @ID deletes bank account
// @Accept json
// @Produce json
// @Param input body inputAccount true "account type"
// @Success 200 {string} string responses.Success
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /account [delete]
func (h *Handler) deleteAccount(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")

	token, err := h.services.Authorization.ParseJwtToken(jwtToken)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var input inputAccount

	err = c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	if h.services.AccountValidator.IsEmptyAccountType(input.Type) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.EmptyAccountType,
		})
	}

	if h.services.AccountValidator.IsMainAccountType(input.Type) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteMainAccount,
		})
	}

	userId, err := typeconv.ConvertStringIdToUintId(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	account, err := h.services.AccountFinder.GetAccountByType(userId, input.Type)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.AccountNotFound,
		})
	}

	if !h.services.AccountValidator.IsZeroBalance(account.Balance) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteNonZeroBalanceAccount,
		})
	}

	err = h.services.Account.DeleteAccount(account)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteAccount,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}
