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
	// 初始化 mock service 和 handler
	mockService := new(MockUserService)
	handler := NewHandler(mockService)

	// 初始化路由
	router := gin.Default()
	router.POST("/user", handler.CreateUser)

	t.Run("Create user", func(t *testing.T) {
		// 測試用戶
		user := domain.User{Name: "John Doe", Email: "john@example.com"}
		userJSON, _ := json.Marshal(user)

		// 預期 mock service 的行為
		mockService.On("CreateUser", user).Return(nil)

		// 創建 HTTP 請求
		req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(userJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		// 創建 ResponseRecorder 來記錄回應
		resp := httptest.NewRecorder()

		// 發送 HTTP 請求
		router.ServeHTTP(resp, req)

		// 斷言 HTTP 狀態碼為 200 OK
		assert.Equal(t, http.StatusOK, resp.Code)

		// 斷言 mock service 的期望調用
		mockService.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	mockService := new(MockUserService)
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/user/:id", handler.GetUserById)

	t.Run("Get user By Id", func(t *testing.T) {
		user := domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

		mockService.On("GetUserById", 1).Return(&user, nil)

		req, err := http.NewRequest(http.MethodGet, "/user/1", nil)
		require.NoError(t, err)

		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		// 拿到200之後確認response json資料正不正確
		var response domain.User
		err = json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 1, response.ID)
		assert.Equal(t, "John Doe", response.Name)
		assert.Equal(t, "john@example.com", response.Email)

		mockService.AssertExpectations(t)
	})
}
