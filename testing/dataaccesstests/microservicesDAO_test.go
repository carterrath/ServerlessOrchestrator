package dataaccesstests

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbMicroservice  *gorm.DB
	daoMicroservice *dataaccess.MicroservicesDAO
	testName        string
	testName2       string
	lastID          uint
)

func TestMicroservicesDAOSuite(t *testing.T) {
	// Setup
	dbMicroservice = setupMicroTestDatabase()
	daoMicroservice = dataaccess.NewMicroservicesDAO(dbMicroservice)

	lastID, err := getLastMicroserviceID(dbMicroservice)
	if err != nil {
		t.Fatalf("Failed to get last microservice ID: %v", err)
	}

	testName = "testmicroservice" + strconv.Itoa(int(lastID+1))
	testName2 = "testmicroservice" + strconv.Itoa(int(lastID+2))

	// Run tests
	t.Run("TestMicroservicesDAO_GetAll", TestMicroservicesDAO_GetAll)
	t.Run("TestMicroservicesDAO_Insert", TestMicroservicesDAO_Insert)
	t.Run("TestMicroservicesDAO_GetByID", TestMicroservicesDAO_GetByID)
	t.Run("TestMicroservicesDAO_GetByName", TestMicroservicesDAO_GetByName)
	t.Run("TestMicroservicesDAO_Update", TestMicroservicesDAO_Update)
	t.Run("TestMicroservicesDAO_Delete", TestMicroservicesDAO_Delete)

	// Teardown
	teardownMicroTestDatabase(dbMicroservice)
}

func teardownMicroTestDatabase(db *gorm.DB) {
	// Clean up test data from the database
	db.Exec("DELETE FROM microservices WHERE backend_name LIKE 'testmicroservice%'")
}

func setupMicroTestDatabase() *gorm.DB {
	// Fetch environment variables
	Username := os.Getenv("POSTGRES_USERNAME")
	//Password := os.Getenv("POSTGRES_PASSWORD")
	Password := ""
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

func getLastMicroserviceID(db *gorm.DB) (uint, error) {
	var microservice business.Microservice
	result := db.Order("id desc").First(&microservice)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// No microservice found, return 0, nil
		return 0, nil
	}
	if result.Error != nil {
		return 0, result.Error
	}
	return microservice.ID, nil
}

func TestMicroservicesDAO_GetAll(t *testing.T) {
	// Test
	microservices, err := daoMicroservice.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, microservices)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_Insert(t *testing.T) {
	// Test
	micro := business.Microservice{
		FriendlyName:  "name",
		RepoLink:      "https://github.com/example/repo",
		StatusMessage: "active",
		IsActive:      true,
		UserID:        1,
		Inputs:        nil,
		OutputLink:    "https://output.link",
		BackendName:   testName,
		ImageID:       "imageid",
	}
	err := daoMicroservice.Insert(micro)

	// Assert
	assert.NoError(t, err)
	// Get the newly inserted record
	newMicro, err := daoMicroservice.GetByName(testName)
	assert.NoError(t, err)
	assert.NotNil(t, newMicro)

	lastID = newMicro.ID // Update lastID to the ID of the newly inserted record
}

func TestMicroservicesDAO_Delete(t *testing.T) {
	// Test
	err := daoMicroservice.Delete(lastID)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_GetByID(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByID(lastID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_GetByName(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByName(testName)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_Update(t *testing.T) {
	// Test
	micro := business.Microservice{
		// create a test microservice object
		FriendlyName:  "name",
		RepoLink:      "https://github.com/example/repo",
		StatusMessage: "active",
		IsActive:      true,
		UserID:        1,
		Inputs:        []business.Input{{Name: "input1", DataType: "string"}},
		OutputLink:    "https://output.link",
		BackendName:   testName2,
		ImageID:       "imageid",
	}
	err := daoMicroservice.Update(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}
