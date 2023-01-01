package responses

const (
	Success           = "success"
	UserAlreadyExists = "user with this username already exists"
	EmptyUsername     = "username cannot be empty"
	EmptyPassword     = "password cannot be empty"

	AccountAlreadyExists           = "account with this type already exists"
	EmptyAccountType               = "account type cannot be empty"
	Unauthorized                   = "unauthorized"
	IncorrectUsernameAndOrPassword = "incorrect username and/or password"

	FailedCreateAccount               = "cannot create account"
	FailedDeleteMainAccount           = "cannot delete main account"
	FailedDeleteAccount               = "cannot delete account"
	FailedDeleteNonZeroBalanceAccount = "cannot delete account, because there are funds on it"
	FailedToParseBody                 = "invalid input data"
	FailedToGenerateJwtToken          = "cannot generate jwt token"

	UserNotfound          = "user not found"
	SenderAccountNotFound = "sender's account not found"
	PayeeAccountNotFound  = "payee's account not found"
	AccountNotFound       = "account not found"
	AccountsNotFound      = "accounts not found "
	TransfersNotFound     = "transfers not found"
	DepositsNotFound      = "deposits not found"

	SenderIsPayee            = "cannot make transfer to your own account with this input transfer type"
	InsufficientFunds        = "insufficient funds to make a transfer"
	CannotTransferZeroAmount = "amount of transfer should be more than 0"
	UnsupportedTransferType  = "unsupported transfer type"

	CannotMakeDatabaseUpdate = "cannot update database value"
)
