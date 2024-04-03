package users

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context, userDAO *dataaccess.UserDAO) {
	// Get the token from the request headers
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		// If the token is missing, return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		return
	}

	// Extract the token from the "Bearer" scheme
	// The token is in the format: "Bearer <token>"
	// We need to extract just the token part
	// Split the string by space and get the second part
	token := strings.Split(tokenString, " ")[1]

	// Get the secret key from the environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		fmt.Println("JWT_SECRET_KEY environment variable not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Parse the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Failed to parse token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		fmt.Println("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Failed to extract claims from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Extract each claim from the claims map
	idFloat64, ok := claims["id"].(float64)
	if !ok {
		fmt.Println("Failed to extract ID from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := int(idFloat64)

	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("Failed to extract username from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	createdAtString, ok := claims["createdAt"].(string)
	if !ok {
		fmt.Println("Failed to extract createdAt from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	createdAt, err := time.Parse(time.RFC3339, createdAtString)
	if err != nil {
		fmt.Println("Failed to parse createdAt:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	updatedAtString, ok := claims["updatedAt"].(string)
	var updatedAt *time.Time
	if ok {
		updatedAtTime, err := time.Parse(time.RFC3339, updatedAtString)
		if err != nil {
			fmt.Println("Failed to parse updatedAt:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		updatedAt = &updatedAtTime
	}

	deletedAtString, ok := claims["deletedAt"].(string)
	var deletedAt *time.Time
	if ok {
		deletedAtTime, err := time.Parse(time.RFC3339, deletedAtString)
		if err != nil {
			fmt.Println("Failed to parse deletedAt:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		deletedAt = &deletedAtTime
	}

	email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("Failed to extract email from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userType, ok := claims["userType"].(string)
	if !ok {
		fmt.Println("Failed to extract userType from token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Return the user details
	c.JSON(http.StatusOK, gin.H{
		"isLoggedIn": true,
		"id":         id,
		"username":   username,
		"createdAt":  createdAt,
		"updatedAt":  updatedAt,
		"deletedAt":  deletedAt,
		"email":      email,
		"userType":   userType,
	})
}
