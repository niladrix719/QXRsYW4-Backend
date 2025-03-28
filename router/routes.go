package router

import (
	"fmt"
	"net/http"

	"QXRsYW4-Backend/controller"
	"QXRsYW4-Backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the application routes
func SetupRouter(r *gin.Engine) {
	// health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Auth routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controller.HandleLogin)
		auth.POST("/register", controller.HandleRegister)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthenticateJWT())
	{
		api.GET("/protected", func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s! This is a protected route.", username)})
		})
	}
}
