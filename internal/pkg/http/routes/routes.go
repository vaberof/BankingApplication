package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/signup", controllers.Register)
	app.Post("/auth", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/balance", controllers.Balance)
	app.Post("/account", controllers.CreateAccount)
	app.Delete("/account", controllers.DeleteAccount)
	app.Post("/transfer", controllers.Transfer)
	app.Get("/history", controllers.Transactions)
}
