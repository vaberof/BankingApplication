package user

type UserStorage interface {
	CreateUser(username string, password string) (*User, error)
	GetUserById(userId uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
}
