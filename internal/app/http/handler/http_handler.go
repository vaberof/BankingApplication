package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vaberof/banking_app/docs"
)

type Handler struct {
	userService    UserService
	accountService AccountService
	authService    AuthorizationService
}

func NewHandler(userService UserService, accountService AccountService, authService AuthorizationService) *Handler {
	return &Handler{
		userService:    userService,
		accountService: accountService,
		authService:    authService,
	}
}

func (h *Handler) InitRoutes(config *fiber.Config) *fiber.App {
	app := fiber.New(*config)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/signup", h.signup)
	app.Post("/login", h.login)
	//app.Post("/logout", h.logout)
	//app.Get("/balance", h.getBalance)
	app.Post("/account", h.createAccount)
	//app.Delete("/account", h.deleteAccount)
	//app.Post("/transfer", h.transfer)
	//app.Get("/transfers", h.getTransfers)
	//app.Get("/deposits", h.getDeposits)

	return app
}
