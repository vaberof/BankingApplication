package user

type AccountStorage interface {
	CreateInitialAccount(userId uint) error
}
