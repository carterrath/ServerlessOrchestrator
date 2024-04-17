package services

import (
	"fmt"

	"github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func StopService(backendNameStr string, dao *dataaccess.MicroservicesDAO) error {
	// get microservice object from database
	microservice, err := FetchMicroservice(backendNameStr, dao)
	if err != nil {
		return fmt.Errorf("error fetching microservice: %w", err)
	}

	// stop the microservice
	err = StopImage(microservice.BackendName)
	if err != nil {
		return fmt.Errorf("error stopping image: %w", err)
	}

	microservice.IsActive = false
	microservice.OutputLink = ""
	// Update the microservice record in the database
	err = dao.Update(*microservice)
	if err != nil {
		return fmt.Errorf("error updating microservice: %w", err)
	}

	return nil
}

func StopImage(backendName string) error {
	// stop the image
	err := dockerhub.StopImage(backendName)
	if err != nil {
		return fmt.Errorf("error stopping image: %w", err)
	}

	return nil
}
