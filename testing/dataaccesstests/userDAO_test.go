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
	dbUser   *gorm.DB
	daoUser  *dataaccess.UserDAO
	testUser *business.User
)

func TestUserDAOSuite(t *testing.T) {
	// Setup
	dbUser = setupTestDatabase()
	daoUser = dataaccess.NewUserDAO(dbUser)
	lastID, err := getLastUserID(dbUser)
	if err != nil {
		t.Fatalf("Failed to get last user ID: %v", err)
	}

	username := "testuser" + strconv.Itoa(lastID+1)
	email := "test" + strconv.Itoa(lastID+1) + "@example.com"

	testUser = &business.User{
		Username: username,
		Email:    email,
		Password: "password",
	}

	// Run tests
	t.Run("TestUserDAO_CreateUser", TestUserDAO_CreateUser)
	t.Run("TestUserDAO_GetUserByUsername", TestUserDAO_GetUserByUsername)
	t.Run("TestUserDAO_GetUserByEmail", TestUserDAO_GetUserByEmail)
	t.Run("TestUserDAO_CheckUsernameAndPassword", TestUserDAO_CheckUsernameAndPassword)
}

func setupTestDatabase() *gorm.DB {
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
	dbUser, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return dbUser
}

func getLastUserID(db *gorm.DB) (int, error) {
	var user business.User
	result := db.Order("id desc").First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, result.Error
	}
	return int(user.ID), nil
}

func TestUserDAO_CreateUser(t *testing.T) {
	// Test
	err := daoUser.CreateUser(testUser)

	// Assert
	assert.NoError(t, err)
}

func TestUserDAO_GetUserByUsername(t *testing.T) {
	// Test
	user, err := daoUser.GetUserByUsername(testUser.Username)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserDAO_GetUserByEmail(t *testing.T) {
	// Test
	user, err := daoUser.GetUserByEmail(testUser.Email)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserDAO_CheckUsernameAndPassword(t *testing.T) {
	// Test
	user, err := daoUser.CheckUsernameAndPassword(testUser.Username, testUser.Password)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
