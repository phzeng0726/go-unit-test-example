package service

import (
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUser(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("GetUser", 1).Return(&domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)

	service := NewUserService(mockRepo)
	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)

	mockRepo.AssertExpectations(t)
}
