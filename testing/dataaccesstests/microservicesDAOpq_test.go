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
	db  *gorm.DB
	dao *dataaccess.MicroservicesDAOpq
)

func TestMicroservicesDAOpqSuite(t *testing.T) {
	// Setup
	db = setupTestDatabase()
	dao = dataaccess.NewMicroservicesDAO(db)

	// Run tests
	t.Run("TestMicroservicesDAOpq_GetAll", TestMicroservicesDAOpq_GetAll)
	t.Run("TestMicroservicesDAOpq_Insert", TestMicroservicesDAOpq_Insert)
	t.Run("TestMicroservicesDAOpq_Delete", TestMicroservicesDAOpq_Delete)
	t.Run("TestMicroservicesDAOpq_GetByID", TestMicroservicesDAOpq_GetByID)
	t.Run("TestMicroservicesDAOpq_GetByName", TestMicroservicesDAOpq_GetByName)
	t.Run("TestMicroservicesDAOpq_Update", TestMicroservicesDAOpq_Update)
}

func setupTestDatabase() *gorm.DB {
	// Fetch environment variables
	Username := os.Getenv("POSTGRES_USERNAME")
	//Password := os.Getenv("POSTGRES_PASSWORD")
	Password := ""
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_DB")

	// Construct the data source name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DB, Port)

	// Open a GORM database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}

func TestMicroservicesDAOpq_GetAll(t *testing.T) {
	// Test
	microservices, err := dao.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, microservices)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Insert(t *testing.T) {
	// Test
	micro := business.Microservice{
		// create a test microservice object
	}
	err := dao.Insert(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Delete(t *testing.T) {
	// Test
	err := dao.Delete(1)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByID(t *testing.T) {
	// Test
	micro, err := dao.GetByID(1)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByName(t *testing.T) {
	// Test
	micro, err := dao.GetByName("test")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Update(t *testing.T) {
	// Test
	micro := business.Microservice{
		// create a test microservice object
	}
	err := dao.Update(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}
