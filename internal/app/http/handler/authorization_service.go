package handler

import (
	"github.com/gofiber/fiber/v2"
	getuser "github.com/vaberof/banking_app/internal/service/user"
)

type AuthorizationService interface {
	AuthenticateUser(jwtToken string) (*getuser.GetUser, error)
	GenerateJwtToken(username string, password string) (string, error)
	GenerateCookie(token string) *fiber.Cookie
	RemoveCookie() *fiber.Cookie
}
