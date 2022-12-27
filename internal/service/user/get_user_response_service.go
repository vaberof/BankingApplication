package user

import (
	domain "github.com/vaberof/MockBankingApplication/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type GetUserResponseService struct {
	userStorage UserStorage
}

func NewGetUserService(userStorage UserStorage) *GetUserResponseService {
	return &GetUserResponseService{userStorage: userStorage}
}

func (s *GetUserResponseService) GetUser(username string, password string) (*GetUserResponse, error) {
	return s.getUserImpl(username, password)
}

func (s *GetUserResponseService) GetUserById(userId uint) (*GetUserResponse, error) {
	return s.getUserByIdImpl(userId)
}

func (s *GetUserResponseService) getUserImpl(username string, password string) (*GetUserResponse, error) {
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

func (s *GetUserResponseService) getUserByIdImpl(userId uint) (*GetUserResponse, error) {
	domainUser, err := s.userStorage.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	getUser := s.domainUserToGetUser(domainUser)
	return getUser, nil
}

func (s *GetUserResponseService) domainUserToGetUser(domainUser *domain.User) *GetUserResponse {
	var getUser GetUserResponse

	getUser.Id = domainUser.Id

	return &getUser
}

func (s *GetUserResponseService) validatePassword(hashedPassword string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
