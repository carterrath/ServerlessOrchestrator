package services

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func SaveMicroservice(microservice business.Microservice, dao *dataaccess.MicroservicesDAO) error {

	// Validate Github URL
	if err := ValidateGithubURL(microservice.RepoLink); err != nil {
		return fmt.Errorf("invalid repository link: %w", err)
	}

	// Generate backend name
	microservice.BackendName = GenerateBackendName(microservice.RepoLink)

	// call CheckIfExists to MicroservicesDAO
	exists, err := CheckIfExists(microservice.BackendName, dao)
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
		return fmt.Errorf("the directory does not contain both Dockerfile and YAML file")
	}

	// call BuildImage, should return image ID
	//BuildImage()
	// return error to api if build fails
	// get user ID from userDAO
	//GetUserID()
	// add image ID and user ID to microservice struct

	// call Insert to MicroservicesDAO
	//Insert(microservice)
	// return error to api if insert fails

	// delete cloned repo from the directory

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

func CheckIfExists(name string, dao *dataaccess.MicroservicesDAO) (bool, error) {
	_, err := dao.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no microservice is found, return false indicating it's safe to proceed
			return false, nil
		}
		// If there's another error querying the database, return it
		return false, fmt.Errorf("error querying database for name %s: %w", name, err)
	}
	// If a microservice is found, return true to indicate it exists
	return true, nil
}

func CloneRepo(repoLink, backendName string) error {
	err := dataaccess.CloneRepositoryUsingCommand(repoLink, backendName)
	if err != nil {
		return err
	}
	return nil
}

func CheckConfigs() (bool, error) {
	var dockerfileExists, yamlFileExists bool
	destinationPath := "application/microholder"
	// Walk through the directory and check for Dockerfile and YAML file
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

		// Check if it's a YAML file
		if filepath.Ext(info.Name()) == ".yaml" || filepath.Ext(info.Name()) == ".yml" {
			yamlFileExists = true
		}

		return nil
	})

	if err != nil {
		return false, err
	}

	// Check if both Dockerfile and YAML file exist
	if err != nil || !dockerfileExists || !yamlFileExists {
		fmt.Println("Dockerfile or YAML file not found or error occurred during directory check:", err)
		fmt.Println("Deleting the directory:", destinationPath)
		if removeErr := os.RemoveAll(destinationPath); removeErr != nil {
			fmt.Println("Error deleting directory:", removeErr)
		}
		return false, err
	}

	return true, nil
}

func BuildImage() (string, error) {
	// call BuildImage to DockerAPI
	// return image ID
	return "", nil
}

func GetUserID() (uint, error) {
	// call GetUserID to UserDAO
	// return user ID
	return 0, nil
}

func Insert(microservice business.Microservice) error {
	// call Insert to MicroservicesDAO
	// return error if fails
	return nil
}
