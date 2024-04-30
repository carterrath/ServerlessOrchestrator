package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/GoKubes/ServerlessOrchestrator/application/elasticcontainerservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"gorm.io/gorm"
)

func ExecuteService(backendNameStr string, dao *dataaccess.MicroservicesDAO, ecsClient *ecs.Client, r53Client *route53.Client) error {
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

	activeMicroserviceCount, err := dao.GetAllActiveCount()
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error getting active microservice count: %w", err)
	}

	port := 3000 + activeMicroserviceCount
	//if the date is not more recent, then run the image
	err = RunImage(microservice.ImageID, microservice.BackendName, port, ecsClient, r53Client)
	if err != nil {
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error running image: %w", err)
	}

	microservice.IsActive = true

	mode := os.Getenv("DEPLOYMENT_MODE")
	if mode == "cloud" {
		microservice.OutputLink = "https://" + microservice.BackendName + ".serverlessorchestrator.com"
	} else {
		microservice.OutputLink = "http://127.0.0.1:" + strconv.Itoa(port)
	}
	microservice.StatusMessage = "Active"

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

func RunImage(imageID, backendName string, port int, ecsClient *ecs.Client, r53Client *route53.Client) error {
	mode := os.Getenv("DEPLOYMENT_MODE")
	domain := os.Getenv("DOMAIN")
	if mode == "local" {
		// Run the image locally on the specified port
		err := dockerhub.RunImageFromDockerHub(imageID, backendName, port)
		if err != nil {
			return err
		}
	} else if mode == "cloud" {
		// Step 1: Register Task Definition
		taskDefinitionArn, err := elasticcontainerservice.RegisterTaskDefinition(ecsClient, backendName)
		if err != nil {
			return fmt.Errorf("error registering task definition: %w", err)
		}

		// Step 2: Create ECS Service
		serviceName := backendName + "-service" // You might want to create a naming convention for services
		err = elasticcontainerservice.CreateService(ecsClient, "MicroserviceOrchestratorCluster", serviceName, *taskDefinitionArn)
		if err != nil {
			return fmt.Errorf("error creating service: %w", err)
		}

		loadBalancerDNSName := os.Getenv("LOAD_BALANCER_DNS_NAME")

		// Step 3: Create DNS Record
		err = elasticcontainerservice.CreateDNSRecord(r53Client, domain, backendName, loadBalancerDNSName)
		if err != nil {
			return fmt.Errorf("error creating DNS record: %w", err)
		}
	}
	return nil
}
