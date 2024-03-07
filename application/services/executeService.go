package services

import (
	"errors"
	"fmt"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"gorm.io/gorm"
)

func ExecuteMicroservice(backendNameStr string, dao *dataaccess.MicroservicesDAO) error {
	// get microservice object from database
	microservice, err := FetchMicroservice(backendNameStr, dao)
	if err != nil {
		return fmt.Errorf("error fetching microservice: %w", err)
	}
	println("Microservice: ", microservice.BackendName)
	// check if image is the latest update of repo
	//get repoLink from Microservice object
	//get date of latest commit to github, if the date is more recent than the updated date on the microservice then delete amd rebuild image.
	//if the date is not more recent, then run the image

	// run image
	return nil
}

// Assume Microservice is a struct that represents your microservice data model
func FetchMicroservice(name string, microserviceDao *dataaccess.MicroservicesDAO) (*business.Microservice, error) {
	microservice, err := microserviceDao.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no microservice is found, return nil and a not found error
			return nil, fmt.Errorf("microservice with name %s not found", name)
		}
		// If there's another error querying the database, return it
		return nil, fmt.Errorf("error querying database for name %s: %w", name, err)
	}
	// Return the found microservice object
	return &microservice, nil
}
