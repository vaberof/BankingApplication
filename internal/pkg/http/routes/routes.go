package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/signup", controllers.Signup)
	app.Post("/auth", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/balance", controllers.GetBalance)
	app.Post("/account", controllers.CreateAccount)
	app.Delete("/account", controllers.DeleteAccount)
	app.Post("/transfer", controllers.MakeTransfer)
	app.Get("/transfers", controllers.GetTransfers)
	app.Get("/deposits", controllers.GetDeposits)
}
