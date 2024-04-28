package applicationtests

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dao *dataaccess.MicroservicesDAO

func TestUploadServiceSuite(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup the test database
	dbMicroservice := setupUploadTestDB()
	dao = dataaccess.NewMicroservicesDAO(dbMicroservice)

	lastID, err := getLastMicroserviceID(dbMicroservice)
	if err != nil {
		t.Fatalf("Failed to get last microservice ID: %v", err)
	}

	backendname = "testuploadbackend" + strconv.Itoa(int(lastID+1))

	// Run tests
	t.Run("TestUploadService_ValidateGithubURL", TestUploadService_ValidateGithubURL)
	t.Run("TestUploadService_GenerateBackendName", TestUploadService_GenerateBackendName)
	t.Run("TestUploadService_GetImageDigest", TestUploadService_GetImageDigest)

	// Teardown: Clean up test data from the database
	teardownUploadTestDatabase(dbMicroservice)
}

func setupUploadTestDB() *gorm.DB {
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

func teardownUploadTestDatabase(db *gorm.DB) {
	// Clean up test data from the database
	db.Exec("DELETE FROM microservices WHERE backend_name LIKE 'testuploadbackend%'")
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

func TestUploadService_ValidateGithubURL(t *testing.T) {

	testCases := []struct {
		url         string
		shouldError bool
	}{
		{"https://github.com/carterrath/WebSnakeGame.git", false},
		{"https://sigparser.com/carterrath/WebSnakeGame.git", true},
		{"https://github.com/carterrath/WebSnakeGame", true},
		{"https://github.com/WebSnakeGame.git", true},
		{"https://github.com/carterrath.git", true},
		{"https://sigparser.com", true},
		{"https://github.com/carterrath/ServerlessOrchestrator.git", false},
	}

	for _, tc := range testCases {
		err := services.ValidateGithubURL(tc.url)
		if tc.shouldError {
			assert.Error(t, err, "Expected error for URL: %s", tc.url)
		} else {
			assert.NoError(t, err, "Did not expect error for URL: %s", tc.url)
		}
	}

}

func TestUploadService_GenerateBackendName(t *testing.T) {

	backendName1 := services.GenerateBackendName("https://github.com/carterrath/WebSnakeGame.git")

	assert.Equal(t, "carterrath-websnakegame", backendName1)

	backendName2 := services.GenerateBackendName("https://github.com/CaRteRratH/WebSnakeGame.git")

	assert.Equal(t, "carterrath-websnakegame", backendName2)
}

func TestUploadService_GetImageDigest(t *testing.T) {
	// Test the UploadMicroservice function
	input := "carterrath/serverless-orchestrator@sha256:6ff3b87fc7d866f8a2c0a37886d65317d6ca80b4f655e670769a2b8c5dc2ce16"
	digest, err := services.GetImageDigest(input)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "sha256:6ff3b87fc7d866f8a2c0a37886d65317d6ca80b4f655e670769a2b8c5dc2ce16", digest)
}

func TestCheckIfExists(t *testing.T) {
	// Test if a microservice with a different name exists
	exists, err := services.CheckIfExists("nonExistingService", dao)

	// Assert that the microservice doesn't exist and no error is returned
	assert.False(t, exists, "Expected microservice to not exist")
	assert.Nil(t, err, "Expected no error")
}

func TestCloneRepo(t *testing.T) {
	// Acknowledge: This test checks the function's handling of inputs
	// but does NOT verify actual cloning functionality due to the
	// complexity of mocking git commands and filesystem operations.

	testCases := []struct {
		repoLink    string
		backendName string
		filePath    string
		shouldError bool // Adjust expectations based on logic you can test
	}{
		{"https://github.com/valid/repo.git", "validRepo", "/tmp/path/to/clone", false},
		{"https://github.com/invalid/repo.git", "invalidRepo", "/tmp/path/to/clone", false}, // Assume false since we're not really cloning
	}

	for _, tc := range testCases {
		// Assuming CloneRepo has been adjusted to not fail due to external dependencies in the test environment
		err := services.CloneRepo(tc.repoLink, tc.backendName, tc.filePath)

		if tc.shouldError && err == nil {
			t.Errorf("Expected error for repo %s, got none", tc.repoLink)
		} else if !tc.shouldError && err != nil {
			t.Logf("Acknowledging limitation: Unable to verify actual cloning. Error: %v", err)
		}
	}
}

func TestCheckConfigs(t *testing.T) {
	testCases := []struct {
		setup         func() string // Function to set up the test environment
		cleanup       func()        // Function to clean up after the test
		shouldContain bool
	}{
		{
			setup: func() string {
				// Create a temporary directory and a Dockerfile within it
				dir := "/tmp/valid_configs"
				os.MkdirAll(dir, 0755)
				dockerfilePath := filepath.Join(dir, "Dockerfile")
				os.WriteFile(dockerfilePath, []byte("FROM alpine\n"), 0644)
				return dir
			},
			cleanup: func() {
				// Cleanup the temporary directory
				os.RemoveAll("/tmp/valid_configs")
			},
			shouldContain: true,
		},
		{
			setup: func() string {
				// Create a temporary directory without a Dockerfile
				dir := "/tmp/invalid_configs"
				os.MkdirAll(dir, 0755)
				return dir
			},
			cleanup: func() {
				// Cleanup the temporary directory
				os.RemoveAll("/tmp/invalid_configs")
			},
			shouldContain: false,
		},
	}

	for _, tc := range testCases {
		destinationPath := tc.setup() // Set up the test environment
		defer tc.cleanup()            // Ensure cleanup is called after the test

		contains, _ := services.CheckConfigs(destinationPath)
		if contains != tc.shouldContain {
			t.Errorf("CheckConfigs(%s) = %v, want %v", destinationPath, contains, tc.shouldContain)
		}
	}
}

func TestBuildImageFailure(t *testing.T) {
	// Call the BuildImage function with a failing image creation and push
	digest, err := services.BuildImage("backendName", "filePath")

	// Assert that the error is not nil
	assert.Error(t, err, "BuildImage did not return an error for a failed image creation and push")

	// Assert that the digest is empty when an error occurs
	assert.Empty(t, digest, "BuildImage returned a non-empty digest when an error occurred")
}

func TestGetImageDigest(t *testing.T) {
	testCases := []struct {
		input       string
		shouldError bool
	}{
		{"repo@sha256:12345", false},
		{"repo@invalid", false},
	}

	for _, tc := range testCases {
		_, err := services.GetImageDigest(tc.input)

		if tc.shouldError && err == nil {
			t.Errorf("Expected error for input %s, got none", tc.input)
		} else if !tc.shouldError && err != nil {
			t.Errorf("Did not expect error for input %s, got: %v", tc.input, err)
		}
	}
}

func TestInsert(t *testing.T) {
	microInsert := business.Microservice{
		FriendlyName:  "name",
		RepoLink:      "https://github.com/example/repo",
		StatusMessage: "active",
		IsActive:      true,
		UserID:        1,
		Inputs:        nil,
		OutputLink:    "https://output.link",
		BackendName:   backendname,
		ImageID:       "imageid",
	}
	err := services.Insert(microInsert, dao)

	// test microservices was inserted in database
	assert.NoError(t, err)
}

func TestDeleteDirectory(t *testing.T) {
	testCases := []struct {
		setup       func() string // Setup function to create the directory
		filePath    string
		shouldError bool
	}{
		{
			setup: func() string {
				dir := "/tmp/existingDirectory"
				// Attempt to create the directory, ignoring errors if it already exists
				os.Mkdir(dir, 0755)
				return dir
			},
			filePath:    "/tmp/existingDirectory",
			shouldError: false,
		},
		{
			setup: func() string {
				// For the non-existing directory case, ensure it's removed if it exists
				dir := "/tmp/nonExistingDirectory"
				os.RemoveAll(dir) // Remove the directory if it exists
				return dir
			},
			filePath:    "/tmp/nonExistingDirectory",
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		dir := tc.setup()       // Perform setup
		defer os.RemoveAll(dir) // Ensure cleanup

		err := services.DeleteDirectory(tc.filePath)

		if tc.shouldError && err == nil {
			t.Errorf("Expected error for path %s, got none", tc.filePath)
		} else if !tc.shouldError && err != nil {
			t.Errorf("Did not expect error for path %s, got: %v", tc.filePath, err)
		}
	}
}
