// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"QXRsYW4-Backend/middleware"
	"QXRsYW4-Backend/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Ensure JWT_SECRET is set
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
		return
	}

	r := gin.Default()

	// Use CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Setup routes
	router.SetupRouter(r)

	// Log that server is starting
	fmt.Println("Server starting on http://0.0.0.0:8080")
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
