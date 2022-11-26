package handler

type inputAccount struct {
	Name string `json:"name"`
}

/*func (h *Handler) deleteAccount(c *fiber.Ctx) error {
	jwtToken := c.Cookies("jwt")

	token, err := h.services.Authorization.ParseJwtToken(jwtToken)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": responses.Unauthorized,
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var input inputAccount

	err = c.BodyParser(&input)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedToParseBody,
		})
	}

	if h.services.AccountValidator.IsEmptyAccountType(input.Type) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.EmptyAccountType,
		})
	}

	if h.services.AccountValidator.IsMainAccountType(input.Type) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteMainAccount,
		})
	}

	userId, err := typeconv.ConvertStringIdToUintId(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	account, err := h.services.AccountFinder.GetAccountByType(userId, input.Type)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": responses.AccountNotFound,
		})
	}

	if !h.services.AccountValidator.IsZeroBalance(account.Balance) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteNonZeroBalanceAccount,
		})
	}

	err = h.services.Account.DeleteAccount(account)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": responses.FailedDeleteAccount,
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": responses.Success,
	})
}*/
