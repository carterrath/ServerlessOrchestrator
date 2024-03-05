package users

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
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

	// Check the username and password against the database
	user, err := userDAO.CheckUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		// If there's an error (user not found or password mismatch), return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
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

	// Additional authentication checks can be added here, google?

	// Return success response with user details
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
