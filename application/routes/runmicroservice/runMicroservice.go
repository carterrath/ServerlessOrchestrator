package runmicroservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

// Define a variable to hold the input string
type BackendNameDto struct {
	Value string `json:"value"`
}

func RunMicroservice(c *gin.Context, dao *dataaccess.MicroservicesDAO) {
	var input BackendNameDto

	// Bind the incoming JSON to the variable
	if err := c.BindJSON(&input); err != nil {
		// Handle error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call executeMicroservice with the string
	err := services.ExecuteMicroservice(input.Value, dao)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, gin.H{"result": "Microservice executed successfully"})
}
