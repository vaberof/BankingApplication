package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(config fiber.Config) *fiber.App {
	app := fiber.New(config)

	app.Post("/signup", h.signUp)
	app.Post("/auth", h.login)
	app.Post("/logout", h.logout)
	app.Get("/balance", h.getBalance)
	app.Post("/account", h.createAccount)
	app.Delete("/account", h.deleteAccount)
	app.Post("/transfer", h.transfer)
	app.Get("/transfers", h.getTransfers)
	app.Get("/deposits", h.getDeposits)

	return app
}
