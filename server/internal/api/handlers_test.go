package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phzeng0726/go-unit-test-example/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	mockService.On("GetUser", 1).Return(&domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}, nil)

	handler := NewHandler(mockService)

	r := gin.Default()
	r.GET("/user/:id", handler.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response domain.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, response.ID)
	assert.Equal(t, "John Doe", response.Name)
	assert.Equal(t, "john@example.com", response.Email)

	mockService.AssertExpectations(t)
}
