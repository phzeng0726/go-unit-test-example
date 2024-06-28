package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite

	mockService *MockUserService
	handler     *Handler
	router      *gin.Engine
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserById(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func (suite *UserHandlerTestSuite) SetupTest() {
	suite.mockService = new(MockUserService)
	suite.handler = NewHandler(suite.mockService)
	suite.router = gin.Default()
}

func (suite *UserHandlerTestSuite) TestCreateUser() {
	suite.router.POST("/user", suite.handler.CreateUser)

	// Test user
	user := domain.User{Name: "Rita Zeng", Email: "rita.zeng@example.com"}
	userJSON, _ := json.Marshal(user)

	// Expected behavior of mock service
	suite.mockService.On("CreateUser", user).Return(nil)

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
	suite.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	// Create ResponseRecorder to record the response
	resp := httptest.NewRecorder()

	// Send HTTP request
	suite.router.ServeHTTP(resp, req)

	// Assert HTTP status code is 200 OK
	suite.Equal(http.StatusOK, resp.Code)

	// Assert expected calls to mock service
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestGetUserById() {

	suite.router.GET("/user/:id", suite.handler.GetUserById)

	user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

	suite.mockService.On("GetUserById", 1).Return(&user, nil)

	req, err := http.NewRequest(http.MethodGet, "/user/1", nil)
	suite.NoError(err)

	resp := httptest.NewRecorder()

	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)

	// After getting 200, verify if the response JSON data is correct
	var response domain.User
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal(1, response.ID)
	suite.Equal("Rita Zeng", response.Name)
	suite.Equal("rita.zeng@example.com", response.Email)

	suite.mockService.AssertExpectations(suite.T())
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
