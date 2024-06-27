package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phzeng0726/go-unit-test-example/internal/api"
	"github.com/phzeng0726/go-unit-test-example/internal/repository"
	"github.com/phzeng0726/go-unit-test-example/internal/service"
)

func main() {
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)
	handler := api.NewHandler(userService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/user/:id", handler.GetUser)

	fmt.Println("Listening 8080...")
	r.Run(":8080")
}
