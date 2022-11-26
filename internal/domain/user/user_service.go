package user

import (
	infra "github.com/vaberof/banking_app/internal/infra/storage/postgres/userpg"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{userStorage: userStorage}
}

func (s *UserService) CreateUser(username string, password string) error {
	return s.createUserImpl(username, password)
}

func (s *UserService) GetUser(username string) (*User, error) {
	return s.getUserImpl(username)
}

func (s *UserService) GetUserById(userId uint) (*User, error) {
	return s.getUserByIdImpl(userId)
}

func (s *UserService) createUserImpl(username string, password string) error {
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	err = s.userStorage.CreateUser(username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) getUserImpl(username string) (*User, error) {
	infraUser, err := s.userStorage.GetUser(username)
	if err != nil {
		return nil, err
	}

	domainUser := s.infraUserToDomain(infraUser)

	return domainUser, nil
}

func (s *UserService) getUserByIdImpl(userId uint) (*User, error) {
	infraUser, err := s.userStorage.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	domainUser := s.infraUserToDomain(infraUser)

	return domainUser, nil
}

func (s *UserService) infraUserToDomain(infraUser *infra.User) *User {
	var user User

	user.Id = infraUser.Id
	user.Username = infraUser.Username
	user.Password = infraUser.Password

	return &user
}

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
