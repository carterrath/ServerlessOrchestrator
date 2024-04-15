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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbMicroservice  *gorm.DB
	daoMicroservice *dataaccess.MicroservicesDAO
)

func testExecuteServiceSuite(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	fmt.Println("passed")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup
	dbMicroservice = setupExecTestDatabase()
	daoMicroservice = dataaccess.NewMicroservicesDAO(dbMicroservice)

	// Run tests
	t.Run("TestExecuteService_ErrorConditions", TestExecuteService_ErrorConditions)

	// Teardown
	teardownMicroTestDatabase(dbMicroservice)
}

func TestExecuteService_ErrorConditions(t *testing.T) {
	// Setup
	dbMicroservice := setupExecTestDatabase()
	daoMicroservice := dataaccess.NewMicroservicesDAO(dbMicroservice)

	// Create a microservice object with a repository link that doesn't exist
	microservice := business.Microservice{
		FriendlyName:  "testname",
		RepoLink:      "https://github.com/example/nonexistentrepo", // non-existent repository
		StatusMessage: "active",
		IsActive:      true,
		UserID:        1,
		Inputs:        nil,
		OutputLink:    "https://output.link",
		BackendName:   "backendname",
		ImageID:       "imageid",
	}

	// Insert the microservice object into the database
	err := daoMicroservice.Insert(&microservice)
	if err != nil {
		t.Fatalf("Failed to insert microservice: %v", err)
	}

	// Run ExecuteService with the microservice object, which should fail due to non-existent repository
	err = services.ExecuteService("testmicroserviceexec", daoMicroservice)
	if err == nil {
		t.Error("ExecuteService did not return an error for a non-existent repository")
	}

	// Teardown
	teardownMicroTestDatabase(dbMicroservice)
}

func teardownMicroTestDatabase(db *gorm.DB) {
	// Clean up test data from the database
	db.Exec("DELETE FROM microservices WHERE backend_name LIKE 'testmicroserviceexec'")
}

func setupExecTestDatabase() *gorm.DB {
	// Fetch environment variables
	Username := os.Getenv("POSTGRES_USERNAME")
	Password := os.Getenv("POSTGRES_PASSWORD")
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_TEST_DB")

	// Construct the data source name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DB, Port)

	// Open a GORM database connection
	dbMicroExec, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return dbMicroExec
}
