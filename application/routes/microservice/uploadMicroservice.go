package microservice

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

// MicroserviceDto represents the DTO (Data Transfer Object) for creating a microservice
type MicroserviceDto struct {
	FriendlyName string     `json:"FriendlyName"`
	RepoLink     string     `json:"RepoLink"`
	Inputs       []InputDto `json:"Inputs"`
	UserID       uint       `json:"UserID"`
}

// InputDto represents the DTO for input data
type InputDto struct {
	ID             uint   `json:"ID"` // Assuming this is optional or generated by the backend
	MicroserviceID uint   `json:"MicroserviceID"`
	Name           string `json:"Name"`
	DataType       string `json:"DataType"`
}

func UploadMicroservice(c *gin.Context, dao *dataaccess.MicroservicesDAO) {
	var microserviceDto MicroserviceDto

	// Bind JSON request body to the MicroserviceDto struct
	if err := c.BindJSON(&microserviceDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert MicroserviceDto to Microservice entity model
	microservice := MapDtoToEntity(microserviceDto)

	// Attempt to save the microservice
	if err := services.SaveMicroservice(microservice, dao); err != nil {
		// Return error in a consistent format
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert microservice into the database
	// if err := dao.Insert(microservice); err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert microservice into database"})
	//     return
	// }

	// Return success message in a consistent format
	c.JSON(http.StatusCreated, gin.H{"message": "Microservice created successfully"})
}

// MapDtoToEntity converts MicroserviceDto to Microservice entity model
func MapDtoToEntity(dto MicroserviceDto) business.Microservice {
	var inputs []business.Input

	// Map InputDto slice to business.Input slice
	for _, inputDto := range dto.Inputs {
		inputs = append(inputs, business.Input{
			Name:     inputDto.Name,
			DataType: inputDto.DataType,
		})
	}

	return business.Microservice{
		FriendlyName: dto.FriendlyName,
		RepoLink:     dto.RepoLink,
		UserID:       dto.UserID,
		Inputs:       inputs, // Assign the mapped inputs slice
	}
}

func MapInputDtoToEntity(dto InputDto, microserviceID uint) business.Input {
	return business.Input{
		Name:     dto.Name,
		DataType: dto.DataType,
	}
}
