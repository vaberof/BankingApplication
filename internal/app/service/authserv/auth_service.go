package authserv

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/MockBankingApplication/internal/app/domain"
	"github.com/vaberof/MockBankingApplication/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user *domain.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateJwtToken(userId uint) (string, error) {
	tokenWithClaims := generateJwtClaims(userId)

	secretKey := os.Getenv("secret_key")
	token, err := tokenWithClaims.SignedString([]byte(secretKey))
	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthService) ParseJwtToken(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthService) GenerateCookie(token string) *fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	return &cookie
}

func (s *AuthService) RemoveCookie() *fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	return &cookie
}

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword)
}

func generateJwtClaims(userId uint) *jwt.Token {
	tokenWithClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(userId)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		})

	return tokenWithClaims
}
