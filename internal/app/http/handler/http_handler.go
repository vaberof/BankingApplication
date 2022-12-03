package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/vaberof/MockBankingApplication/docs"
)

type HttpHandler struct {
	userService    UserService
	accountService AccountService
	authService    AuthorizationService
}

func NewHttpHandler(userService UserService, accountService AccountService, authService AuthorizationService) *HttpHandler {
	return &HttpHandler{
		userService:    userService,
		accountService: accountService,
		authService:    authService,
	}
}

func (h *HttpHandler) InitRoutes(config *fiber.Config) *fiber.App {
	app := fiber.New(*config)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/register", h.register)
	app.Post("/login", h.login)
	app.Post("/logout", h.logout)

	app.Post("/account", h.createAccount)
	app.Delete("/account", h.deleteAccount)
	app.Get("/accounts", h.getAccounts)

	//app.Post("/transfer", h.transfer)
	//app.Get("/transfers", h.getTransfers)
	//app.Get("/deposits", h.getDeposits)

	return app
}
