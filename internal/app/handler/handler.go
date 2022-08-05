package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) InitRoutes(config fiber.Config) *fiber.App {
	app := fiber.New(config)

	app.Post("/signup", h.Signup)
	app.Post("/login", h.Login)
	app.Post("/logout", h.Logout)
	app.Get("/balance", h.GetBalance)
	app.Post("/account", h.CreateAccount)
	app.Delete("/account", h.DeleteAccount)
	app.Post("/transfer", h.MakeTransfer)
	app.Get("/transfers", h.GetTransfers)
	app.Get("/deposits", h.GetDeposits)

	return app
}
