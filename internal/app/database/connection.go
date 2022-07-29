package database

import (
	"fmt"
	"github.com/vaberof/banking_app/internal/app/model"
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

	connection.AutoMigrate(model.User{})
	connection.AutoMigrate(model.Account{})
	connection.AutoMigrate(model.Transfer{})
}
