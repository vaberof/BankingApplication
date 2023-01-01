package handler

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/vaberof/MockBankingApplication/internal/domain/user"
)

type AuthenticationService interface {
	AuthenticateUser(jwtToken string) (*domain.User, error)
	GenerateJwtToken(username string, password string) (string, error)
	GenerateCookie(token string) *fiber.Cookie
	RemoveCookie() *fiber.Cookie
}
