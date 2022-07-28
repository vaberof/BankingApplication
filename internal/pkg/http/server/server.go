package server

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func Run(app *fiber.App) {
	app.Listen(os.Getenv("server_host") + ":" + os.Getenv("server_port"))
}
