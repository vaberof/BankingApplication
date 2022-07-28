package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func CreateUser(inputUsername string, hashedPassword []byte) *model.User {
	user := model.NewUser()

	user.SetUsername(inputUsername)
	user.SetPassword(hashedPassword)

	return user
}

func GetUser(data map[string]string) (*model.User, error) {
	user := model.NewUser()

	result := database.DB.Table("users").Where("username = ?", data["username"]).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func FindUserById(userID string) (*model.User, error) {
	user := model.NewUser()

	result := database.DB.Table("users").Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func CreateUserInDatabase(user *model.User) {
	database.DB.Create(&user)
}

func IsCorrectPassword(hashedPassword []byte, password string) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return false
	}
	return true
}

func HashPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return hashedPassword
}

func CreateJwtClaims(user *model.User) *jwt.Token {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		})

	return claims
}

func CreateJwtToken(claims *jwt.Token) (string, error) {
	secretKey := os.Getenv("secret_key")
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return token, err
	}

	return token, nil
}

func CreateCookie(token string) *fiber.Cookie {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	}

	return &cookie
}

func RemoveCookie() *fiber.Cookie {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	return &cookie
}

func ParseJwtToken(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("secret_key")
		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
