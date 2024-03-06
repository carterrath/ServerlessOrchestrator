package microservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

type MicroserviceData struct {
	ID           uint
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	FriendlyName string
	RepoLink     string
	Status       string
	User         business.User
	Inputs       []business.Input
	OutputLink   string
	BackendName  string `gorm:"unique"`
}

func GetAllMicroservices(c *gin.Context, dao *dataaccess.MicroservicesDAO, userDAO *dataaccess.UserDAO) {
	microservices, err := dao.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	microservicesData := GetAllMicroservicesData(microservices, userDAO)
	c.IndentedJSON(http.StatusOK, microservicesData)
}

func GetAllMicroservicesData(microservices []business.Microservice, userDAO *dataaccess.UserDAO) []MicroserviceData {
	var microservicesData []MicroserviceData

	for _, microservice := range microservices {
		userInterface, err := userDAO.GetByID(microservice.UserID)
		if err != nil {
			continue
		}

		// Assuming GetByID is supposed to return a *business.User
		user, ok := userInterface.(*business.User)
		if !ok {
			continue
		}

		microservicesData = append(microservicesData, MicroserviceData{
			ID:           microservice.ID,
			CreatedAt:    microservice.CreatedAt.String(),
			UpdatedAt:    microservice.UpdatedAt.String(),
			DeletedAt:    microservice.DeletedAt.Time.String(),
			FriendlyName: microservice.FriendlyName,
			RepoLink:     microservice.RepoLink,
			Status:       microservice.Status,
			User:         *user,
			Inputs:       microservice.Inputs,
			OutputLink:   microservice.OutputLink,
			BackendName:  microservice.BackendName,
		})
	}

	return microservicesData
}
