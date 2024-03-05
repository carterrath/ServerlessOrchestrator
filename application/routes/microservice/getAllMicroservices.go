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
	microservicesData := []MicroserviceData{}

	for i := range microservices {
		userId := microservices[i].UserID
		user, err := userDAO.GetUserByID(userId)

		if err != nil {
			microservicesData[i] = MicroserviceData{
				ID:           microservices[i].ID,
				CreatedAt:    microservices[i].CreatedAt.String(),
				UpdatedAt:    microservices[i].UpdatedAt.String(),
				DeletedAt:    microservices[i].DeletedAt.Time.String(),
				FriendlyName: microservices[i].FriendlyName,
				RepoLink:     microservices[i].RepoLink,
				Status:       microservices[i].Status,
				User:         *user,
				Inputs:       microservices[i].Inputs,
				OutputLink:   microservices[i].OutputLink,
				BackendName:  microservices[i].BackendName,
			}
		}
	}

	return microservicesData
}
