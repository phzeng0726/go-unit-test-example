package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "pong", response["message"])
}

func TestUserRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/john", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Hello john", response["message"])
}

func TestCreateUserRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	userData := `{"name":"john","age":30}`
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(userData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "User created", response["message"])

	user, ok := response["user"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "john", user["name"])
	assert.Equal(t, float64(30), user["age"]) // JSON numbers are float64
}
