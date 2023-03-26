package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userStorage    UserStorage
	accountStorage AccountStorage
}

func NewUserService(userStorage UserStorage, accountStorage AccountStorage) *UserService {
	return &UserService{userStorage: userStorage, accountStorage: accountStorage}
}

func (s *UserService) CreateUser(username string, password string) (*User, error) {
	return s.createUserImpl(username, password)
}

func (s *UserService) GetUser(username string, password string) (*User, error) {
	return s.getUserImpl(username, password)
}

func (s *UserService) GetUserById(userId uint) (*User, error) {
	return s.getUserByIdImpl(userId)
}

func (s *UserService) GetUserByUsername(username string) (*User, error) {
	return s.getUserByUsernameImpl(username)
}

func (s *UserService) createUserImpl(username string, password string) (*User, error) {
	_, err := s.userStorage.GetUserByUsername(username)
	if err == nil {
		return nil, errors.New("user with this username already exist")
	}

	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := s.userStorage.CreateUser(username, hashedPassword)
	if err != nil {
		return nil, err
	}

	err = s.accountStorage.CreateInitialAccount(user.Id)
	if err != nil {
		return nil, fmt.Errorf("cannot create initial account: %s", err.Error())
	}

	return user, nil
}

func (s *UserService) getUserImpl(username string, password string) (*User, error) {
	user, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = s.validatePassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) getUserByIdImpl(userId uint) (*User, error) {
	user, err := s.userStorage.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) getUserByUsernameImpl(username string) (*User, error) {
	user, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *UserService) validatePassword(hashedPassword string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
