package user

import "golang.org/x/crypto/bcrypt"

type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService {
	return &UserService{userStorage: userStorage}
}

func (s *UserService) CreateUser(username string, password string) error {
	return s.createUserImpl(username, password)
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

func (s *UserService) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
