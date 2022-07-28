package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/pkg/http/routes"
	"github.com/vaberof/banking_app/internal/pkg/http/server"
	"log"
)

func main() {
	loadEnvironmentVariables()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.Connect()

	routes.Setup(app)

	server.Run(app)
}

func loadEnvironmentVariables() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
