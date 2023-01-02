package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/domain/user"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

func buildUser(user *user.User) *UserResponse {
	return &UserResponse{
		Id:       user.Id,
		Username: user.Username,
	}
}

func RenderUserResponse(c *fiber.Ctx, user *user.User) error {
	return RenderResponse(c, buildUser(user))
}
