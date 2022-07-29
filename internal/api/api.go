package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/banking_app/internal/app/controllers"
)

func Signup(c *fiber.Ctx) error {
	if err := controllers.Register(c); err != nil {
		return err
	}
	return nil
}

func Login(c *fiber.Ctx) error {
	if err := controllers.Login(c); err != nil {
		return err
	}
	return nil
}

func Logout(c *fiber.Ctx) error {
	if err := controllers.Logout(c); err != nil {
		return err
	}
	return nil
}

func GetBalance(c *fiber.Ctx) error {
	if err := controllers.Balance(c); err != nil {
		return err
	}
	return nil
}

func CreateNewAccount(c *fiber.Ctx) error {
	if err := controllers.CreateAccount(c); err != nil {
		return err
	}
	return nil
}

func DeleteAccount(c *fiber.Ctx) error {
	if err := controllers.DeleteAccount(c); err != nil {
		return err
	}
	return nil
}

func MakeTransfer(c *fiber.Ctx) error {
	if err := controllers.Transfer(c); err != nil {
		return err
	}
	return nil
}

func GetTransfers(c *fiber.Ctx) error {
	if err := controllers.Transfers(c); err != nil {
		return err
	}
	return nil
}

func GetDeposits(c *fiber.Ctx) error {
	if err := controllers.Deposits(c); err != nil {
		return err
	}
	return nil
}
