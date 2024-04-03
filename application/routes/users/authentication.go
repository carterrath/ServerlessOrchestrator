package users

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateJWT(user *business.User) (string, error) {
	// Define the token claims
	claims := jwt.MapClaims{
		"id":        user.ID,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
		"deletedAt": user.DeletedAt,
		"email":     user.Email,
		"username":  user.Username,
		"userType":  user.UserType,
		"exp":       time.Now().Add(time.Hour).Unix(), // Token expiry time (e.g., 24 hours)
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get the secret key from the environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	fmt.Println("secret" + secretKey)
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Middleware function to verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		tokenString := c.GetHeader("Authorization")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Get the secret key from the environment variable
			secretKey := os.Getenv("JWT_SECRET_KEY")
			if secretKey == "" {
				return nil, fmt.Errorf("JWT_SECRET_KEY environment variable not set")
			}
			return []byte(secretKey), nil
		})

		// Check if there's an error or if the token is invalid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Token is valid, continue with the request
		c.Next()
	}
}
