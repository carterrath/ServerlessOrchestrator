package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"gorm.io/gorm"
)

func ExecuteService(backendNameStr string, dao *dataaccess.MicroservicesDAO) error {
	// get microservice object from database
	microservice, err := FetchMicroservice(backendNameStr, dao)
	if err != nil {
		return fmt.Errorf("error fetching microservice: %w", err)
	}

	// check if image is the latest update of repo
	filePath := "application/microholder/"
	dateStr, err := GetLatestPushDate(microservice.RepoLink, microservice.BackendName, filePath)
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error getting latest push date: %w", err)
	}
	println("Date: ", dateStr)

	//get date of latest commit to github, if the date is more recent than the updated date on the microservice then delete amd rebuild image.
	date, err := ParseDate(dateStr)
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error parsing date: %w", err)
	}

	fmt.Println("Date: ", date)
	fmt.Println("UpdatedAt: ", microservice.UpdatedAt)

	// Compare the parsed date with the updatedAt field of the microservice
	if date.After(microservice.UpdatedAt) {
		fmt.Println("The repository has been updated more recently than the microservice. Updating microservice...")
		// Rebuild the image
		digest, err := BuildImage(microservice.BackendName, filePath)
		if err != nil {
			_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
			return fmt.Errorf("error building image: %w", err)
		}
		microservice.ImageID = digest
	}
	//if the date is not more recent, then run the image
	err = RunImage(microservice.ImageID, microservice.BackendName, 3000)
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error running image: %w", err)
	}

	microservice.IsActive = true
	// Update the microservice record in the database
	err = dao.Update(*microservice)
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error updating microservice: %w", err)
	}

	err = DeleteDirectory(filePath + microservice.BackendName)
	if err != nil {
		// If deleting cloned directory fails, log the error
		fmt.Printf("failed to delete directory: %v\n", err)
	}

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

func GetLatestPushDate(repoURL, backendName, filePath string) (string, error) {
	date, err := github.GetLatestPushDate(repoURL, backendName, filePath)
	if err != nil {
		return "", err
	}
	return date, nil
}

func ParseDate(dateStr string) (time.Time, error) {
	// Parse the date string into a time.Time object
	date, err := time.Parse("Mon Jan 2 15:04:05 2006 -0700", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func RunImage(imageID, backendName string, port int) error {
	// Run the image locally on the specified port
	err := dockerhub.RunImageFromDockerHub(imageID, backendName, port)
	if err != nil {
		return err
	}
	return nil
}
