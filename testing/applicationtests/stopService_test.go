package applicationtests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestStopServiceWithError(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	fmt.Println("passed")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup the test database
	dbMicroservice := setupTestDB()
	dao := dataaccess.NewMicroservicesDAO(dbMicroservice)

	// Create a new microservice
	microservice := business.Microservice{
		FriendlyName:  "name",
		RepoLink:      "https://github.com/example/repo",
		StatusMessage: "active",
		IsActive:      true,
		UserID:        1,
		Inputs:        nil,
		OutputLink:    "https://output.link",
		BackendName:   "test",
		ImageID:       "imageid",
	}

	// Insert the microservice into the database
	err = dao.Insert(microservice)
	assert.Nil(t, err)

	// Simulate an error condition by manually deleting the microservice
	err = dbMicroservice.Exec("DELETE FROM microservices WHERE backend_name = ?", "test").Error
	assert.Nil(t, err) // Ensure the deletion was successful

	// Attempt to stop the service
	err = services.StopService("test", dao)
	assert.Error(t, err, "StopService did not return an error when expected")

	// Teardown: Clean up test data from the database
	teardownTestDatabase(dbMicroservice)
}

func setupTestDB() *gorm.DB {
	// Fetch environment variables
	Username := os.Getenv("POSTGRES_USERNAME")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_TEST_DB")

	// Construct the data source name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DB, Port)

	// Open a GORM database connection
	dbMicroservice, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return dbMicroservice
}

func teardownTestDatabase(db *gorm.DB) {
	// Clean up test data from the database
	db.Exec("DELETE FROM microservices WHERE backend_name LIKE 'test'")
}