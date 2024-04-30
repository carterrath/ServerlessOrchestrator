package runmicroservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/gin-gonic/gin"
)

// Define a variable to hold the input string
type BackendNameDto struct {
	Value string `json:"value"`
}

func RunMicroservice(c *gin.Context, dao *dataaccess.MicroservicesDAO, ecsClient *ecs.Client, r53Client *route53.Client) {
	var input BackendNameDto

	// Bind the incoming JSON to the variable
	if err := c.BindJSON(&input); err != nil {
		// Handle error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call executeMicroservice with the string
	err := services.ExecuteService(input.Value, dao, ecsClient, r53Client)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, gin.H{"result": "Microservice executed successfully"})
}
