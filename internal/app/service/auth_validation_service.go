package service

import (
	"github.com/vaberof/banking_app/internal/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthValidationService struct {
	repos repository.AuthorizationValidator
}

func NewAuthValidationService(repos repository.AuthorizationValidator) *AuthValidationService {
	return &AuthValidationService{repos: repos}
}

func (s *AuthValidationService) UserExists(username string) error {
	return s.repos.UserExists(username)
}

func (s *AuthValidationService) IsCorrectPassword(hashedPassword string, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}

func (s *AuthValidationService) IsEmptyUsername(username string) bool {
	return len(username) == 0
}

func (s *AuthValidationService) IsEmptyPassword(password string) bool {
	return len(password) == 0
}
