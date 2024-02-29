package orca

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func SaveMicroservice(microservice business.Microservice, dao *dataaccess.MicroservicesDAOpq) {
	// call CheckIfExists to MicroservicesDAOpq
	CheckIfExists(microservice.Name, dao)
	// return error if exists

	// call CheckAccess to GitHubAPI with Github repo link
	CheckAccess(microservice.RepoLink)
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

	// return success to api

}

func CheckIfExists(name string, dao *dataaccess.MicroservicesDAOpq) error {
	_, err := dao.GetByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no microservice is found, return nil indicating it's safe to proceed
			return nil
		}
		// If there's another error querying the database, return it
		return fmt.Errorf("error querying database for name %s: %w", name, err)
	}
	// If a microservice is found, return an error indicating a microservice with this name already exists
	return errors.New("microservice with the same name already exists")
}

func CheckAccess(repoLink string) error {
	// call CheckAccess to GitHubAPI
	// return error if not public

	return nil
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
