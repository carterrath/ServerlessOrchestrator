package users

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateDeveloperRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateConsumerRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateDeveloper(c *gin.Context, userDAO *dataaccess.UserDAO) {
	var req CreateDeveloperRequest
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

	// Create a new developer user instance
	user := business.User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
		UserType: "Developer",
	}

	// Insert the user into the database
	if err := userDAO.Insert(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create developer account", "details": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Developer account created successfully", "user": user})
}

func CreateConsumer(c *gin.Context, userDAO *dataaccess.UserDAO) {
	var req CreateConsumerRequest
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

	// Create a new consumer user instance
	user := business.User{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
		UserType: "Consumer",
	}

	// Insert the user into the database
	if err := userDAO.Insert(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consumer account"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Consumer account created successfully", "user": user})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
