package server

import (
	"github.com/gofiber/fiber/v2"
)

func Run(host string, port string, app *fiber.App) error {
	return app.Listen(host + port)
}
