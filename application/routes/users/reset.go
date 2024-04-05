package users

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ResetPassword(c *gin.Context, userDAO *dataaccess.UserDAO) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Update the user's password in the database
	err = userDAO.UpdatePassword(req.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
