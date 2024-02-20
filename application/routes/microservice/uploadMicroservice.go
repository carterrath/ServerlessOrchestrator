package microservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

type CreateMicroserviceRequest struct {
	Name        string `json:"name" binding:"required"`
	PlaceHolder string `json:"placeholder"`
	Input       string `json:"input"`
}

func UploadMicroservice(c *gin.Context, dao *dataaccess.MicroservicesDAOpq) {
	var req CreateMicroserviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if microservice with the given name already exists
	_, err := dao.GetByName(req.Name)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Microservice with the same name already exists"})
		return
	}

	// Create Microservice instance
	microservice := business.Microservice{
		Name:        req.Name,
		PlaceHolder: req.PlaceHolder,
	}

	// Insert the microservice into the database
	if err := dao.Insert(microservice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert microservice"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Microservice created successfully", "microservice": microservice})

}
