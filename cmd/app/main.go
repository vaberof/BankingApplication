package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vaberof/MockBankingApplication/internal/app/http/handler"
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
	"github.com/vaberof/MockBankingApplication/internal/domain/user"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/accountpg"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/depositpg"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/transferpg"
	"github.com/vaberof/MockBankingApplication/internal/infra/storage/postgres/userpg"
	getaccount "github.com/vaberof/MockBankingApplication/internal/service/account"
	"github.com/vaberof/MockBankingApplication/internal/service/auth"
	"github.com/vaberof/MockBankingApplication/internal/service/deposit"
	"github.com/vaberof/MockBankingApplication/internal/service/transfer"
	getuser "github.com/vaberof/MockBankingApplication/internal/service/user"
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
		User:     os.Getenv("db_username"),
		Password: os.Getenv("db_password"),
	})
	if err != nil {
		log.Fatalf("cannot connect to database %s", err.Error())
	}

	err = db.AutoMigrate(&accountpg.Account{}, &userpg.User{}, &transferpg.Transfer{}, &depositpg.Deposit{})
	if err != nil {
		log.Fatalf("cannot auto migrate models %s", err.Error())
	}

	userStoragePostgres := userpg.NewPostgresUserStorage(db)
	accountStoragePostgres := accountpg.NewPostgresAccountStorage(db)
	transferStoragePostgres := transferpg.NewPostgresTransferStorage(db, accountStoragePostgres)
	depositStoragePostgres := depositpg.NewPostgresDepositStorage(db)

	userService := user.NewUserService(userStoragePostgres, accountStoragePostgres)
	accountService := account.NewAccountService(accountStoragePostgres)
	depositService := deposit.NewDepositService(depositStoragePostgres)

	getUserService := getuser.NewGetUserService(userService)
	getAccountResponseService := getaccount.NewGetAccountService(accountService)

	transferService := transfer.NewTransferService(transferStoragePostgres, depositService, accountStoragePostgres, getUserService)

	authService := auth.NewAuthService(getUserService)

	httpHandler := handler.NewHttpHandler(userService, getAccountResponseService, transferService, authService)

	app := httpHandler.InitRoutes(&fiber.Config{
		WriteTimeout: time.Duration(viper.GetInt("server.write_timeout")) * time.Second,
		ReadTimeout:  time.Duration(viper.GetInt("server.read_timeout")) * time.Second,
	})

	if err = app.Listen(viper.GetString("server.host") + ":" + viper.GetString("server.port")); err != nil {
		log.Fatalf("cannot run a server: %v", err)
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
