package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	domain "github.com/vaberof/banking_app/internal/domain/user"
	getuser "github.com/vaberof/banking_app/internal/service/user"
	"os"
	"strconv"
	"time"
)

type AuthService struct {
	getUserService GetUserService
}

func NewAuthService(getUserService GetUserService) *AuthService {
	return &AuthService{
		getUserService: getUserService,
	}
}

func (s *AuthService) AuthenticateUser(jwtToken string) (*getuser.GetUser, error) {
	token, err := s.parseJwtToken(jwtToken)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	userId, err := domain.FromString(claims.Issuer)
	if err != nil {
		return nil, err
	}

	getUser, err := s.getUserService.GetUserById(uint(userId))
	if err != nil {
		return nil, err
	}

	return getUser, nil
}

func (s *AuthService) GenerateJwtToken(username string, password string) (string, error) {
	tokenWithClaims, err := s.generateJwtClaims(username, password)
	if err != nil {
		return "", err
	}

	token, err := tokenWithClaims.SignedString([]byte(os.Getenv("secret_key")))
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

func (s *AuthService) parseJwtToken(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthService) generateJwtClaims(username string, password string) (*jwt.Token, error) {
	user, err := s.getUserService.GetUser(username, password)
	if err != nil {
		return nil, errors.New("incorrect username or/and password")
	}

	tokenWithClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		})

	return tokenWithClaims, nil
}
