package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// @Summary		Get all bank accounts
// @Tags			Bank Account
// @Description	Get all bank accounts you have
// @ID				gets all bank accounts
// @Produce		json
// @Success		200	{array}	account.GetAccountResponse
// @Failure		401	{object}	error
// @Failure		500	{object}	error
// @Router			/accounts [get]
func (h *HttpHandler) getAccounts(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	log.Printf("COOKIES: %v, error in GET ACCOUNTS: %v", c.Cookies("jwt"), err)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	accounts, err := h.accountService.GetAccounts(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"accounts": accounts,
	})
}
