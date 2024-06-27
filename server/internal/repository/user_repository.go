package repository

import (
	"errors"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserById(id int) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user domain.User) error {
	if user.ID == 0 && user.Name == "" && user.Email == "" {
		return errors.New("cannot create empty user")
	}

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserById(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
