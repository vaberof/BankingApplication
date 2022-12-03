package user

import (
	domain "github.com/vaberof/banking_app/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type GetUserService struct {
	userStorage UserStorage
}

func NewGetUserService(userStorage UserStorage) *GetUserService {
	return &GetUserService{userStorage: userStorage}
}

func (s *GetUserService) CreateUser(username string, password string) error {
	return s.userStorage.CreateUser(username, password)
}

func (s *GetUserService) GetUser(username string, password string) (*GetUser, error) {
	return s.getUserImpl(username, password)
}

func (s *GetUserService) GetUserById(userId uint) (*GetUser, error) {
	return s.getUserByIdImpl(userId)
}

func (s *GetUserService) GetUserByUsername(username string) (*GetUser, error) {
	return s.getUserByUsernameImpl(username)
}

func (s *GetUserService) getUserImpl(username string, password string) (*GetUser, error) {
	domainUser, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = s.validatePassword(domainUser.Password, password)
	if err != nil {
		return nil, err
	}

	getUser := s.domainUserToGetUser(domainUser)
	return getUser, nil
}

func (s *GetUserService) getUserByIdImpl(userId uint) (*GetUser, error) {
	domainUser, err := s.userStorage.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	getUser := s.domainUserToGetUser(domainUser)
	return getUser, nil
}

func (s *GetUserService) getUserByUsernameImpl(username string) (*GetUser, error) {
	domainUser, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	getUser := s.domainUserToGetUser(domainUser)
	return getUser, nil
}

func (s *GetUserService) domainUserToGetUser(domainUser *domain.User) *GetUser {
	var getUser GetUser

	getUser.Id = domainUser.Id
	getUser.Username = domainUser.Username
	getUser.Password = domainUser.Password

	return &getUser
}

func (s *GetUserService) validatePassword(hashedPassword string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
