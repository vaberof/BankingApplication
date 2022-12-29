package handler

import "github.com/gofiber/fiber/v2"

//	@Summary		Get all transfers
//	@Tags			Transfer
//	@Description	Get all transfers you have made
//	@ID				Gets transfers
//	@Produce		json
//	@Success		200	{string}	error
//	@Failure		401	{object}	error
//	@Failure		500	{object}	error
//	@Router			/transfers [get]
func (h *HttpHandler) getTransfers(c *fiber.Ctx) error {
	user, err := h.authService.AuthenticateUser(c.Cookies("jwt"))
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	transfers, err := h.transferService.GetTransfers(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"transfers": transfers,
	})
}
