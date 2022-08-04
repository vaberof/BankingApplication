package database

import (
	"fmt"
	"github.com/vaberof/banking_app/internal/app/domain/account"
	"github.com/vaberof/banking_app/internal/app/domain/deposit"
	"github.com/vaberof/banking_app/internal/app/domain/transfer"
	"github.com/vaberof/banking_app/internal/app/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprint("postgres://" + os.Getenv("db_user") +
		":" + os.Getenv("db_password") +
		"@" + os.Getenv("db_host") +
		":" + os.Getenv("db_port") +
		"/" + os.Getenv("db_name"))

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(user.User{})
	connection.AutoMigrate(account.Account{})
	connection.AutoMigrate(transfer.Transfer{})
	connection.AutoMigrate(deposit.Deposit{})
}
