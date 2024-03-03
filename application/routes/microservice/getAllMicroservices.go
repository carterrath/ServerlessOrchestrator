package microservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

func GetAllMicroservices(c *gin.Context, dao *dataaccess.MicroservicesDAO) {
	microservices, err := dao.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range microservices {
		microservices[i].Inputs, err = dao.GetInputs(microservices[i].ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, microservices)
}
