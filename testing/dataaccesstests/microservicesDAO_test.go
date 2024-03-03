package dataaccesstests

import (
	"fmt"
	"os"
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
)

func TestMicroservicesDAOSuite(t *testing.T) {
	// Setup
	dbMicroservice = setupMicroTestDatabase()
	daoMicroservice = dataaccess.NewMicroservicesDAO(dbMicroservice)

	// Run tests
	t.Run("TestMicroservicesDAO_GetAll", TestMicroservicesDAO_GetAll)
	t.Run("TestMicroservicesDAO_Insert", TestMicroservicesDAO_Insert)
	t.Run("TestMicroservicesDAO_Delete", TestMicroservicesDAO_Delete)
	t.Run("TestMicroservicesDAO_GetByID", TestMicroservicesDAO_GetByID)
	t.Run("TestMicroservicesDAO_GetByName", TestMicroservicesDAO_GetByName)
	t.Run("TestMicroservicesDAO_Update", TestMicroservicesDAO_Update)
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
		// create a test microservice object
	}
	err := daoMicroservice.Insert(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_Delete(t *testing.T) {
	// Test
	err := daoMicroservice.Delete(1)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_GetByID(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByID(1)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_GetByName(t *testing.T) {
	// Test
	micro, err := daoMicroservice.GetByName("test")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAO_Update(t *testing.T) {
	// Test
	micro := business.Microservice{
		// create a test microservice object
	}
	err := daoMicroservice.Update(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}
