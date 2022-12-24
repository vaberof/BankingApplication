package handler

type UserService interface {
	CreateUser(username string, password string) (uint, error)
}
