package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{}

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Update this with your Vue.js frontend URL
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Auth routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handleLogin)
		auth.POST("/register", handleRegister)
	}

	r.Run(":8080")
}

func handleLogin(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	for _, user := range users {
		if user.Username == loginReq.Username && user.Password == loginReq.Password {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"user":    user.Username,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func handleRegister(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if username already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    newUser.Username,
	})
}
