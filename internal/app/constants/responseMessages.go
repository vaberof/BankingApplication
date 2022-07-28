package constants

const (
	Success                        = "success"
	UserAlreadyExists              = "user with this username already exists"
	AccountAlreadyExists           = "account with this type already exists"
	EmptyAccountType               = "account type cannot be empty"
	Unauthorized                   = "user unauthorized"
	IncorrectUsernameAndOrPassword = "incorrect username and/or password"

	FailedLogin    = "could not login"
	FailedTransfer = "transaction is not possible"

	SenderAccountNotFound = "sender's account not found"
	PayeeAccountNotFound  = "payee's account not found"
	AccountsNotFound      = "accounts not found "
	TransactionsNotFound  = "transactions not found"

	SenderIsRecipient = "you are trying to make a transfer to your own account"
	InsufficientFunds = "insufficient funds"

	UnsupportedTransferType   = "unsupported transfer type"
	UnsupportedTransferAmount = "unsupported transfer amount"
)
