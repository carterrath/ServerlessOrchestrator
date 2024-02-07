package dataaccesstests

import (
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMicroservicesDAOpq_GetAll(t *testing.T) {
	// Setup
	db = // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	// Test
	microservices, err := dao.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, microservices)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Insert(t *testing.T) {
	// Setup
	db := // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	micro := business.Microservice{
		// create a test microservice object
	}

	// Test
	err := dao.Insert(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_Delete(t *testing.T) {
	// Setup
	db := // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	// Test
	err := dao.Delete(1)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByID(t *testing.T) {
	// Setup
	db := // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	// Test
	micro, err := dao.GetByID(1)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestMicroservicesDAOpq_GetByName(t *testing.T) {
	// Setup
	db := // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	// Test
	micro, err := dao.GetByName("test")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, micro)
	// Add more assertions based on your requirements
}

func TestNicroservicesDAOpq_Update(t *testing.T) {
	// Setup
	db := // create a test database connection

	dao := dataaccess.NewMicroservicesDAO(db)

	micro := business.Microservice{
		// create a test microservice object
	}

	// Test
	err := dao.Update(micro)

	// Assert
	assert.NoError(t, err)
	// Add more assertions based on your requirements
}