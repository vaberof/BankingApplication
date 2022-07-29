package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/api"
)

func Setup(app *fiber.App) {
	app.Post("/signup", api.Signup)
	app.Post("/auth", api.Login)
	app.Post("/logout", api.Logout)
	app.Get("/balance", api.GetBalance)
	app.Post("/account", api.CreateNewAccount)
	app.Delete("/account", api.DeleteAccount)
	app.Post("/transfer", api.MakeTransfer)
	app.Get("/transfers", api.GetTransfers)
	app.Get("/deposits", api.GetDeposits)
}
