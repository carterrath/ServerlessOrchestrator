package orca

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func SaveMicroservice(microservice business.Microservice, dao *dataaccess.MicroservicesDAOpq) error {
	// call CheckIfExists to MicroservicesDAOpq
	exists, err := CheckIfExists(microservice.Name, dao)
	if err != nil {
		return fmt.Errorf("failed to check if microservice exists: %w", err)
	}
	if exists {
		return errors.New("microservice with the same name already exists")
	}

	// call CloneRepo
	if err := CloneRepo(microservice.RepoLink); err != nil {
		return fmt.Errorf("failed to clone repository: %w, Make sure the repository is public", err)
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

	// call Insert to MicroservicesDAOpq
	//Insert(microservice)
	// return error to api if insert fails

	// delete cloned repo from the directory

	// return success to api
	return nil
}

func CheckIfExists(name string, dao *dataaccess.MicroservicesDAOpq) (bool, error) {
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

func CloneRepo(repoLink string) error {
	err := dataaccess.CloneRepository(repoLink)
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
	// call GetUserID to UserDAOpq
	// return user ID
	return 0, nil
}

func Insert(microservice business.Microservice) error {
	// call Insert to MicroservicesDAOpq
	// return error if fails
	return nil
}
