package orca

import (
	"errors"
	"fmt"

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
	// return error if exists

	// call CheckAccess to GitHubAPI with Github repo link
	isPublic, err := CheckAccess(microservice.RepoLink)
	if err != nil {
		return fmt.Errorf("failed to check repo access: %w", err)
	}
	if !isPublic {
		return errors.New("repository is private")
	}
	// return error to api if repo link is not public

	// call CloneRepo
	//CloneRepo(microservice.RepoLink)
	// return any error to api

	// call CheckConfigs
	//CheckConfigs()
	// return error to api if missing dockerfile, yaml file, buildscript

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

func CheckAccess(repoLink string) (bool, error) {
	// call IsPublicRepo to check if the repository is public
	isPublic, err := dataaccess.IsPublicRepo(repoLink)
	if err != nil {
		return false, err // Return the error immediately if IsPublicRepo returns an error
	}
	if !isPublic {
		return false, nil
	}
	return true, nil
}

func CloneRepo(repoLink string) error {
	// call CloneRepo to GitHubAPI
	// return error if fails
	return nil
}

func CheckConfigs() error {
	// call CheckConfigs to GitHubAPI
	// return error if missing
	return nil
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
