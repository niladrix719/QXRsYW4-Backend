package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret = func() []byte {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return []byte(secret)
	}
	return []byte("secret")
}()

// CORSMiddleware configures CORS settings
func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://vue-service:5173",
	}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	return cors.New(config)
}

// GenerateJWT generates a new JWT token for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Token expires in 1 hour
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("username", claims.Username) // Store username in context
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		}
	}
}
