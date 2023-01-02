package views

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vaberof/MockBankingApplication/internal/domain/account"
)

type AccountResponse struct {
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func buildAccount(account *account.Account) *AccountResponse {
	return &AccountResponse{
		Id:      account.Id,
		Type:    account.Type,
		Name:    account.Name,
		Balance: account.Balance,
	}
}

func RenderAccountResponse(c *fiber.Ctx, account *account.Account) error {
	return RenderResponse(c, buildAccount(account))
}

func RenderAccountsResponse(c *fiber.Ctx, accounts []*account.Account) error {
	accountsResponse := make([]*AccountResponse, len(accounts))

	for i := 0; i < len(accountsResponse); i++ {
		accountsResponse[i] = buildAccount(accounts[i])
	}

	return RenderResponse(c, accountsResponse)
}
