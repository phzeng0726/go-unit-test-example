package service

import (
	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/phzeng0726/go-unit-test-example/internal/repository"
)

type UserService interface {
	CreateUser(user domain.User) error
	GetUserById(id int) (*domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserById(id int) (*domain.User, error) {
	return s.repo.GetUserById(id)
}
