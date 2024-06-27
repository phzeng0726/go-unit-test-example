package service

import (
	"errors"
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) GetUserById(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	s := NewUserService(mockRepo)

	t.Run("Create user", func(t *testing.T) {
		user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}
		mockRepo.On("CreateUser", user).Return(nil)
		err := s.CreateUser(user)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create user with empty data", func(t *testing.T) {
		emptyUser := domain.User{}
		mockRepo.On("CreateUser", emptyUser).Return(errors.New("cannot create empty user"))
		err := s.CreateUser(emptyUser)
		assert.Error(t, err)
		assert.Equal(t, "cannot create empty user", err.Error())
		mockRepo.AssertExpectations(t)
	})

}

func TestGetUserById(t *testing.T) {
	mockRepo := new(MockUserRepo)
	s := NewUserService(mockRepo)

	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

	t.Run("Get existing user by id", func(t *testing.T) {
		mockRepo.On("GetUserById", user.ID).Return(&user, nil)
		user, err := s.GetUserById(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "Rita Zeng", user.Name)
		assert.Equal(t, "rita.zeng@example.com", user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get non-existing user by id", func(t *testing.T) {
		nonExistingID := 2
		mockRepo.On("GetUserById", nonExistingID).Return(nil, errors.New("user not found"))
		_, err := s.GetUserById(nonExistingID)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
