package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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

func TestCreateUser(t *testing.T) {
	// Initialize mock service and handler
	mockService := new(MockUserService)
	handler := NewHandler(mockService)

	// Initialize router
	router := gin.Default()
	router.POST("/user", handler.CreateUser)

	t.Run("Create user", func(t *testing.T) {
		// Test user
		user := domain.User{Name: "Rita Zeng", Email: "rita.zeng@example.com"}
		userJSON, _ := json.Marshal(user)

		// Expected behavior of mock service
		mockService.On("CreateUser", user).Return(nil)

		// Create HTTP request
		req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		// Create ResponseRecorder to record the response
		resp := httptest.NewRecorder()

		// Send HTTP request
		router.ServeHTTP(resp, req)

		// Assert HTTP status code is 200 OK
		assert.Equal(t, http.StatusOK, resp.Code)

		// Assert expected calls to mock service
		mockService.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/user/:id", handler.GetUserById)

	t.Run("Get user By Id", func(t *testing.T) {
		user := domain.User{ID: 1, Name: "Rita Zeng", Email: "rita.zeng@example.com"}

		mockService.On("GetUserById", 1).Return(&user, nil)

		req, err := http.NewRequest(http.MethodGet, "/user/1", nil)
		require.NoError(t, err)

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		// After getting 200, verify if the response JSON data is correct
		var response domain.User
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, "Rita Zeng", response.Name)
		assert.Equal(t, "rita.zeng@example.com", response.Email)

		mockService.AssertExpectations(t)
	})
}
