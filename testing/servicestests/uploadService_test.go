package servicestests

import (
	"os"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/stretchr/testify/assert"
)

func TestUploadServiceSuite(t *testing.T) {

	// Run tests
	t.Run("TestUploadService_ValidateGithubURL", TestUploadService_ValidateGithubURL)
	t.Run("TestUploadService_GenerateBackendName", TestUploadService_GenerateBackendName)
	t.Run("TestUploadService_TestDeleteDirectory", TestUploadService_TestDeleteDirectory)
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

	backendName := services.GenerateBackendName("https://github.com/carterrath/WebSnakeGame.git")

	assert.Equal(t, "carterrath-WebSnakeGame", backendName)
}

func TestUploadService_TestDeleteDirectory(t *testing.T) {
	github.CloneRepositoryUsingCommand("https://github.com/carterrath/WebSnakeGame.git", "carterrath-WebSnakeGame")
	err := services.DeleteDirectory("/Users/carterrath/Documents/Fall2023/SE490/ServerlessOrchestrator/application/microholder/carterrath-WebSnakeGame")
	if err != nil {
		t.Fatalf("failed to delete directory: %v", err)
	}

	// Check if the directory exists
	_, err = os.Stat("/Users/carterrath/Documents/Fall2023/SE490/ServerlessOrchestrator/application/microholder/carterrath-WebSnakeGame")
	if err == nil {
		// Directory still exists, test failed
		t.Error("directory still exists after deletion")
	}

}
