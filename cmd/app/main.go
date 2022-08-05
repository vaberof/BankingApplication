package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vaberof/banking_app/internal/app/database"
	"github.com/vaberof/banking_app/internal/app/handler"
	"github.com/vaberof/banking_app/internal/pkg/http/server"
	"log"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed initializating configs: %s", err.Error())
	}

	if err := loadEnvironmentVariables(); err != nil {
		log.Fatalf("failed loading environment variables: %s", err.Error())
	}

	config := fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	handlers := new(handler.Handler)
	app := handlers.InitRoutes(config)

	database.Connect()

	if err := server.Run(viper.GetString("server.host"), viper.GetString("server.port"), app); err != nil {
		log.Fatalf("cannot run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func loadEnvironmentVariables() error {
	err := godotenv.Load("../../.env")
	return err
}
