package user

type UserStorage interface {
	CreateUser(username string, password string) error
}
