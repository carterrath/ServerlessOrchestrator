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

func SaveMicroservice(microservice business.Microservice, microserviceDao *dataaccess.MicroservicesDAO, userID uint) error {

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
	filePath := "application/microholder/"
	if err := CloneRepo(microservice.RepoLink, microservice.BackendName, filePath); err != nil {
		// If cloning fails, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("failed to upload microservice: make sure the repository is public: %w", err)
	}

	// call CheckConfigs
	containsFiles, err := CheckConfigs(filePath + microservice.BackendName)
	if err != nil {
		// If error occurs, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error when checking repo: %w", err)
	}

	if !containsFiles {
		// If necessary files are not found, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("the directory does not contain Dockerfile")
	}

	// call BuildImage, should return image ID
	res, err := BuildImage(microservice.BackendName, filePath)
	if err != nil {
		// If building image fails, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("failed to build image: %w", err)
	}
	println("Image ID: ", res)
	digest, err := GetImageDigest(res)
	if err != nil {
		// If getting image digest fails, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("error: %v", err)
	}
	microservice.ImageID = digest

	// return error to api if build fails
	// get user ID from userDAO
	// GetUserID()
	// microservice.UserID = 1

	// Assign the current user's ID to the microservice
	microservice.UserID = userID

	// add image ID and user ID to microservice struct

	// call Insert to MicroservicesDAO
	err = Insert(microservice, microserviceDao)
	if err != nil {
		// If inserting microservice fails, delete the cloned directory
		_ = DeleteDirectory(filePath + microservice.BackendName) // Ignoring error here as we're already returning an error
		return fmt.Errorf("failed to insert microservice: %w", err)
	}
	// return error to api if insert fails

	// delete cloned repo from the directory

	err = DeleteDirectory(filePath + microservice.BackendName)
	if err != nil {
		// If deleting cloned directory fails, log the error
		fmt.Printf("failed to delete directory: %v\n", err)
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
		name := pathComponents[0] + "-" + repoName
		// Convert the name to lowercase before returning
		return strings.ToLower(name)
	}

	// If there are fewer than two components, return the path as is
	return strings.ToLower(repoPath)
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

func CloneRepo(repoLink, backendName, filePath string) error {
	err := github.CloneRepositoryUsingCommand(repoLink, backendName, filePath)
	if err != nil {
		return err
	}
	return nil
}

func CheckConfigs(destinationPath string) (bool, error) {
	var dockerfileExists bool
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
		return false, err
	}

	return true, nil
}

func BuildImage(backendName, filePath string) (string, error) {
	digest, err := dockerhub.CreateAndPushImage(backendName, filePath)
	if err != nil {
		return digest, err
	}
	return digest, nil
}

func GetImageDigest(input string) (string, error) {
	// Split the input string by "@" to separate the repository name and image digest
	parts := strings.Split(input, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format")
	}

	// Return only the image digest part
	return parts[1], nil
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
