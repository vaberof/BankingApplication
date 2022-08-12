package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vaberof/banking_app/internal/app/handler"
	"github.com/vaberof/banking_app/internal/app/repository"
	"github.com/vaberof/banking_app/internal/app/repository/postgres"
	"github.com/vaberof/banking_app/internal/app/service"
	"github.com/vaberof/banking_app/internal/pkg/http/server"
	"log"
	"os"
	"time"
)

// @title Banking App
// @version 1.0
// @description API Server for Mock Banking Application

// @host localhost:8080
// @BasePath /

// @securityDefinition.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed initializating configs: %s", err.Error())
	}

	if err := loadEnvironmentVariables(); err != nil {
		log.Fatalf("failed loading environment variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDb(&postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Name:     viper.GetString("db.name"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("db_password"),
	})
	if err != nil {
		log.Fatalf("cannot connect to database %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := handlers.InitRoutes(fiber.Config{
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	})

	if err = repos.MakeMigrations(db); err != nil {
		log.Fatalf("cannot make migrations %s", err.Error())
	}

	if err = server.Run(viper.GetString("server.host"), viper.GetString("server.port"), app); err != nil {
		log.Fatalf("cannot run server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs_/")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func loadEnvironmentVariables() error {
	err := godotenv.Load("../../.env")
	return err
}
