package users

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"userType" binding:"required"`
}

func Login(c *gin.Context, userDAO *dataaccess.UserDAO) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username exists in the database
	user, err := userDAO.GetUserByUsername(req.Username)
	if err != nil {
		// If there's an error (user not found), return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// Compare the hashed password stored in the database with the entered password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		// If there's an error (password mismatch), return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Check if the user is trying to login with the correct UserType
	if req.UserType == "Consumer" && user.UserType != "Consumer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a Consumer, use Developer Login"})
		return
	} else if req.UserType == "Developer" && user.UserType != "Developer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not a Developer, use Consumer Login"})
		return
	}

	// Generate a JWT token
	token, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token:" + err.Error()})
		return
	}

	// Return success response with user details
	// c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})

	// Return success response with user details and token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
