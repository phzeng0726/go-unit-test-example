package repository

import "github.com/phzeng0726/go-unit-test-example/internal/domain"

type UserRepository interface {
	GetUser(id int) (*domain.User, error)
}

type userRepository struct {
	// 這裡可以是數據庫連接
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUser(id int) (*domain.User, error) {

	// 模擬從數據庫獲取用戶
	return &domain.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}
