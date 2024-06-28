package service

import (
	"errors"
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite

	mockRepo    *MockUserRepo
	userService UserService
}

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

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepo)
	suite.userService = NewUserService(suite.mockRepo)
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}
	suite.mockRepo.On("CreateUser", user).Return(nil)

	err := suite.userService.CreateUser(user)
	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserServiceTestSuite) TestCreateUser_WithEmptyDataError() {
	emptyUser := domain.User{}
	suite.mockRepo.On("CreateUser", emptyUser).Return(errors.New("cannot create empty user"))

	err := suite.userService.CreateUser(emptyUser)
	suite.Error(err)
	suite.Equal("cannot create empty user", err.Error())
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserServiceTestSuite) TestGetUserById() {
	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}
	suite.mockRepo.On("GetUserById", user.ID).Return(&user, nil)

	fetchedUser, err := suite.userService.GetUserById(user.ID)

	suite.NoError(err)
	suite.NotNil(fetchedUser)
	suite.Equal(1, fetchedUser.ID)
	suite.Equal("Rita Zeng", fetchedUser.Name)
	suite.Equal("rita.zeng@example.com", fetchedUser.Email)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserServiceTestSuite) TestGetUserById_WithInvalidError() {
	nonExistingID := 2
	suite.mockRepo.On("GetUserById", nonExistingID).Return(nil, errors.New("user not found"))

	_, err := suite.userService.GetUserById(nonExistingID)

	suite.Error(err)
	suite.Equal("user not found", err.Error())
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
