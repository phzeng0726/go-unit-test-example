package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// main.go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
			"user":    user,
		})
	})
	return r
}

func main() {
	r := setupRouter()

	fmt.Println("Listening 8080...")
	r.Run(":8080")
}
