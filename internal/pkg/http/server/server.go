package server

import (
	"github.com/gofiber/fiber/v2"
)

func Run(port string, app *fiber.App) error {
	return app.Listen(port)
}
