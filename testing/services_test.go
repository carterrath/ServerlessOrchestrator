package testing

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	servicedb       *gorm.DB
	daoServiceMicro *dataaccess.MicroservicesDAO
)

func TestServicesSuite(t *testing.T) {
	/*
		1. Setup the test database
		2. Create a test microservice object (will use an functioning microservice)
		3. use the Microservice to test uploadService.go
		4. use the Microservice to test executeService.go
		5. use the Microservice to test stopService.go
		6. Teardown the test database
	*/
	servicedb = setupServicesTestDatabase()
	createTestMicroservice()

	t.Run("TestUploadService", TestUploadService)
	t.Run("TestExecuteService", TestExecuteService)
	t.Run("TestStopService", TestStopService)

	teardownServicesTestDatabase()
}

func setupServicesTestDatabase() *gorm.DB {
	// Setup the test database

	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	fmt.Println("passed")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetch environment variables
	Username := os.Getenv("POSTGRES_USERNAME")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_TEST_DB")

	// Construct the data source name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DB, Port)

	// Open a GORM database connection
	servicedb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return servicedb
}

func teardownServicesTestDatabase() {
	// Teardown the test database
	servicedb.Exec("DELETE FROM services WHERE FriendlyName LIKE 'testserviceMicro'")
}

func createTestMicroservice() {
	// Create a test microservice object
	serviceMicro := business.Microservice{
		FriendlyName:  "testserviceMicro",
		RepoLink:      "https://github.com/ruthijimenez/shopping-cart.git",
		StatusMessage: "",
		IsActive:      false,
		UserID:        1,
		Inputs:        nil,
		OutputLink:    "",
		BackendName:   "",
		ImageID:       "",
	}

	err := daoServiceMicro.Insert(serviceMicro)
	if err != nil {
		log.Fatalf("Failed to insert test microservice: %v", err)
	}
}

func TestUploadService(t *testing.T) {
	// Test uploadService.go
}

func TestExecuteService(t *testing.T) {
	// Test executeService.go
}

func TestStopService(t *testing.T) {
	// Test stopService.go
}
