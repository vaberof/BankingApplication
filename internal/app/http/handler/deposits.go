package handler

/*
import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/pkg/responses"
	"github.com/vaberof/banking_app/internal/pkg/typeconv"
)

// @Summary Get deposits
// @Tags Deposit
// @Description Get the deposits made to your bank account(s) using personal or client transfers
// @ID gets all deposits
// @Produce json
// @Success 200 {object} domain.Deposits
// @Failure 401 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /deposits [get]
func (h *Handler) getDeposits(c *fiber.Ctx) error {
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

	deposits, err := h.services.Deposit.GetDeposits(userId)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.DepositsNotFound,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"deposits": deposits,
	})
}
*/
