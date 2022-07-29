package constants

const (
	Success                        = "success"
	UserAlreadyExists              = "user with this username already exists"
	AccountAlreadyExists           = "account with this type already exists"
	EmptyAccountType               = "account type cannot be empty"
	Unauthorized                   = "user unauthorized"
	IncorrectUsernameAndOrPassword = "incorrect username and/or password"

	FailedLogin                       = "could not login"
	FailedTransfer                    = "transaction is not possible"
	FailedDeleteMainAccount           = "cannot delete main account"
	FailedDeleteNonZeroBalanceAccount = "cannot delete account, because there are funds on it"

	SenderAccountNotFound = "sender's account not found"
	PayeeAccountNotFound  = "payee's account not found"
	AccountNotFound       = "account not found"
	AccountsNotFound      = "accounts not found "
	TransfersNotFound     = "transfers not found"

	SenderIsRecipient = "you are trying to make a transfer to your own account"
	InsufficientFunds = "insufficient funds"

	UnsupportedTransferType   = "unsupported transfer type"
	UnsupportedTransferAmount = "unsupported transfer amount"
)
