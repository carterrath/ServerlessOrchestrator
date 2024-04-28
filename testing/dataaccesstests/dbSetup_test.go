package dataaccesstests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateDatabase(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Call CreateDatabase function
	db := setupTestDB()

	// Assert that a *gorm.DB object is returned
	assert.IsType(t, &gorm.DB{}, db)

	// AutoMigrate the models
	err = db.AutoMigrate(&business.Input{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&business.Microservice{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&business.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Check if the tables exist in the database
	assert.True(t, db.Migrator().HasTable(&business.Input{}))
	assert.True(t, db.Migrator().HasTable(&business.Microservice{}))
	assert.True(t, db.Migrator().HasTable(&business.User{}))
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
