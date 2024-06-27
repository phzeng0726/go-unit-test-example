package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phzeng0726/go-unit-test-example/internal/api"
	"github.com/phzeng0726/go-unit-test-example/internal/domain"
	"github.com/phzeng0726/go-unit-test-example/internal/repository"
	"github.com/phzeng0726/go-unit-test-example/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	handler := api.NewHandler(userService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetUserById)

	fmt.Println("Listening 8080...")
	r.Run(":8080")
}
