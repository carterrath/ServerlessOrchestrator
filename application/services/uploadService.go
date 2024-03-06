package services

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"

	"github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func SaveMicroservice(microservice business.Microservice, microserviceDao *dataaccess.MicroservicesDAO) error {

	// Validate Github URL
	if err := ValidateGithubURL(microservice.RepoLink); err != nil {
		return fmt.Errorf("invalid repository link: %w", err)
	}

	// Generate backend name
	microservice.BackendName = GenerateBackendName(microservice.RepoLink)

	// call CheckIfExists to MicroservicesDAO
	exists, err := CheckIfExists(microservice.BackendName, microserviceDao)
	if err != nil {
		return fmt.Errorf("failed to check if microservice exists: %w", err)
	}
	if exists {
		return errors.New("microservice with the same name already exists")
	}

	// call CloneRepo
	if err := CloneRepo(microservice.RepoLink, microservice.BackendName); err != nil {
		return fmt.Errorf("failed to upload microservice: make sure the repository is public")
	}

	// call CheckConfigs
	containsFiles, err := CheckConfigs()
	if err != nil {
		return fmt.Errorf("error when checking repo: %w", err)
	}

	if !containsFiles {
		return fmt.Errorf("the directory does not contain Dockerfile")
	}

	// call BuildImage, should return image ID
	digest, err := BuildImage(microservice.BackendName)
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}
	println("Image ID: ", digest)
	microservice.ImageID = digest

	// return error to api if build fails
	// get user ID from userDAO
	//GetUserID()
	microservice.UserID = 4
	// add image ID and user ID to microservice struct

	// call Insert to MicroservicesDAO
	err = Insert(microservice, microserviceDao)
	if err != nil {
		return fmt.Errorf("failed to insert microservice: %w", err)
	}
	// return error to api if insert fails

	// delete cloned repo from the directory
	err = DeleteDirectory("/Users/carterrath/Documents/Fall2023/SE490/ServerlessOrchestrator/application/microholder/" + microservice.BackendName)
	if err != nil {
		return fmt.Errorf("failed to delete directory: %w", err)
	}
	// return success to api
	return nil
}

func ValidateGithubURL(repoLink string) error {
	// Parse the URL
	u, err := url.ParseRequestURI(repoLink)
	if err != nil {
		return fmt.Errorf("the repository link is not a valid URL: %w", err)
	}

	// Check if the host is GitHub
	if !strings.Contains(u.Host, "github.com") {
		return fmt.Errorf("the repository link does not belong to GitHub")
	}

	// Check if the path contains at least two components (username/repo)
	pathComponents := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(pathComponents) < 2 {
		return fmt.Errorf("the repository link is not in the format username/repo")
	}

	// Check if the last component contains ".git"
	repoName := pathComponents[len(pathComponents)-1]
	if !strings.HasSuffix(repoName, ".git") {
		return fmt.Errorf("the repository name should end with .git")
	}

	return nil
}

func GenerateBackendName(repoLink string) string {
	// Parse the repository link URL
	repoURL, err := url.Parse(repoLink)
	if err != nil {
		return "" // Return empty string on error
	}

	// Extract the path component
	repoPath := repoURL.Path

	// Remove the leading slash
	repoPath = strings.TrimPrefix(repoPath, "/")

	// Split the path into components
	pathComponents := strings.Split(repoPath, "/")

	// If there are at least two components (username/repo), return them
	if len(pathComponents) >= 2 {
		// Remove the .git extension from the repository name
		repoName := strings.TrimSuffix(pathComponents[1], ".git")
		return pathComponents[0] + "-" + repoName
	}

	// If there are fewer than two components, return the path as is
	return repoPath
}

func CheckIfExists(name string, microserviceDao *dataaccess.MicroservicesDAO) (bool, error) {
	_, err := microserviceDao.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no microservice is found, return false indicating it's safe to proceed
			return false, nil
		}
		// If there's another error querying the database, return it
		return false, fmt.Errorf("error querying database for name %s: %w", name, err)
	}
	return true, nil
}

func CloneRepo(repoLink, backendName string) error {
	err := github.CloneRepositoryUsingCommand(repoLink, backendName)
	if err != nil {
		return err
	}
	return nil
}

func CheckConfigs() (bool, error) {
	var dockerfileExists bool
	destinationPath := "application/microholder"
	// Walk through the directory and check for Dockerfile
	err := filepath.Walk(destinationPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file
		if !info.Mode().IsRegular() {
			return nil
		}

		// Check if it's a Dockerfile
		if info.Name() == "Dockerfile" {
			dockerfileExists = true
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	// Check if both Dockerfile and YAML file exist
	if err != nil || !dockerfileExists {
		fmt.Println("Dockerfile not found or error occurred during directory check:", err)
		fmt.Println("Deleting the directory:", destinationPath)
		if removeErr := os.RemoveAll(destinationPath); removeErr != nil {
			fmt.Println("Error deleting directory:", removeErr)
		}
		return false, err
	}

	return true, nil
}

func BuildImage(backendName string) (string, error) {
	digest, err := dockerhub.CreateAndPushImage(backendName)
	if err != nil {
		return digest, err
	}
	return digest, nil
}

func GetUserID() (uint, error) {
	// call GetUserID to UserDAO
	// return user ID
	return 0, nil
}

func Insert(microservice business.Microservice, microserviceDao *dataaccess.MicroservicesDAO) error {
	err := microserviceDao.Insert(microservice)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDirectory(filePath string) error {
	// Check if directory exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Directory doesn't exist
			return fmt.Errorf("directory %s does not exist", filePath)
		}
		// Error occurred while checking directory
		return err
	}

	// Attempt to remove directory
	err = os.RemoveAll(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete directory %s: %v", filePath, err)
	}

	return nil
}
