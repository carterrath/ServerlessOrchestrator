package dataaccesstests

import (
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateDatabase(t *testing.T) {
	// Call CreateDatabase function
	db := dataaccess.CreateDatabase()

	// Assert that a *gorm.DB object is returned
	assert.IsType(t, &gorm.DB{}, db)

	// Check if the tables exist in the database
	assert.True(t, db.Migrator().HasTable(&business.Input{}))
	assert.True(t, db.Migrator().HasTable(&business.Microservice{}))
	assert.True(t, db.Migrator().HasTable(&business.User{}))
}
