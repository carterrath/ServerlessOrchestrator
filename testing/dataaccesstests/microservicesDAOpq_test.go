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
	daoMicroservice *dataaccess.MicroservicesDAOpq
	testName        string
	testName2       string
	lastID          uint
)

func TestMicroservicesDAOpqSuite(t *testing.T) {
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
	t.Run("TestMicroservicesDAOpq_GetAll", TestMicroservicesDAOpq_GetAll)
	t.Run("TestMicroservicesDAOpq_Insert", TestMicroservicesDAOpq_Insert)
	t.Run("TestMicroservicesDAOpq_GetByID", TestMicroservicesDAOpq_GetByID)
	t.Run("TestMicroservicesDAOpq_GetByName", TestMicroservicesDAOpq_GetByName)
	t.Run("TestMicroservicesDAOpq_Update", TestMicroservicesDAOpq_Update)
	t.Run("TestMicroservicesDAOpq_Delete", TestMicroservicesDAOpq_Delete)
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
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}
	return microservice.ID, nil
}

func TestMicroservicesDAOpq_GetAll(t *testing.T) {
	// Test
	microservices, err := daoMicroservice.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, microservices)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Insert(t *testing.T) {
	// Test
	micro := business.Microservice{
		Name:       testName,
		RepoLink:   "https://github.com/example/repo",
		Status:     "active",
		Author:     "test_author",
		Inputs:     nil,
		OutputLink: "https://output.link",
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

func TestMicroservicesDAOpq_Delete(t *testing.T) {
	// Test
	err := daoMicroservice.Delete(lastID)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByID(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByID(lastID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByName(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByName(testName)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Update(t *testing.T) {
	// Test
	micro := business.Microservice{
		// create a test microservice object
		Name:       testName2,
		RepoLink:   "https://github.com/example/repo",
		Status:     "active",
		Author:     "test_author",
		Inputs:     []business.Input{{Name: "input1", DataType: "string"}},
		OutputLink: "https://output.link",
	}
	err := daoMicroservice.Update(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}
