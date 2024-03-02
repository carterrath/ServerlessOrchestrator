package microservice

import (
	"net/http"
	"net/url"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/GoKubes/ServerlessOrchestrator/orca"
	"github.com/gin-gonic/gin"
)

// MicroserviceDto represents the DTO (Data Transfer Object) for creating a microservice
type MicroserviceDto struct {
	// ID        uint       `json:"id"` // Assuming this is optional or generated by the backend
	// CreatedAt time.Time  `json:"createdAt"`
	// UpdatedAt time.Time  `json:"updatedAt"`
	// DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Name     string     `json:"name"`
	RepoLink string     `json:"repoLink"`
	Author   string     `json:"author"`
	Inputs   []InputDto `json:"inputs"`
	Status   string     `json:"status"`
}

// InputDto represents the DTO for input data
type InputDto struct {
	ID             uint   `json:"id"` // Assuming this is optional or generated by the backend
	MicroserviceID uint   `json:"microserviceID"`
	Name           string `json:"name"`
	DataType       string `json:"dataType"`
}

func UploadMicroservice(c *gin.Context, dao *dataaccess.MicroservicesDAOpq) {
	var microserviceDto MicroserviceDto

	// Bind JSON request body to the MicroserviceDto struct
	if err := c.BindJSON(&microserviceDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate RepoLink URL
	if _, err := url.ParseRequestURI(microserviceDto.RepoLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid RepoLink URL"})
	}
	// Convert MicroserviceDto to Microservice entity model
	microservice := MapDtoToEntity(microserviceDto)

	// Attempt to save the microservice
	if err := orca.SaveMicroservice(microservice, dao); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert microservice into the database
	// if err := dao.Insert(microservice); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert microservice into database"})
	// 	return
	// }

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
		Name:     dto.Name,
		RepoLink: dto.RepoLink,
		Author:   dto.Author,
		Inputs:   inputs, // Assign the mapped inputs slice
		Status:   dto.Status,
	}
}

func MapInputDtoToEntity(dto InputDto, microserviceID uint) business.Input {
	return business.Input{
		Name:     dto.Name,
		DataType: dto.DataType,
	}
}
