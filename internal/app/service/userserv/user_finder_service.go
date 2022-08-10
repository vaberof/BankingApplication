package userserv

import (
	"github.com/vaberof/banking_app/internal/app/domain"
	"github.com/vaberof/banking_app/internal/app/repository"
)

type UserFinderService struct {
	repos repository.UserFinder
}

func NewUserFinderService(repos repository.UserFinder) *UserFinderService {
	return &UserFinderService{repos: repos}
}

func (s *UserFinderService) GetUserById(userId uint) (*domain.User, error) {
	user, err := s.repos.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserFinderService) GetUserByUsername(username string) (*domain.User, error) {
	user, err := s.repos.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
