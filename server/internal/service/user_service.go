package service

import (
	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/phzeng0726/go-unit-test-example/internal/repository"
)

type UserService interface {
	GetUser(id int) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUser(id int) (*domain.User, error) {
	return s.repo.GetUser(id)
}
