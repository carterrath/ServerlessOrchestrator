package servicestests

import (
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/services"
	"github.com/stretchr/testify/assert"
)

func TestUploadServiceSuite(t *testing.T) {

	// Run tests
	t.Run("TestUploadService_ValidateGithubURL", TestUploadService_ValidateGithubURL)
	t.Run("TestUploadService_GenerateBackendName", TestUploadService_GenerateBackendName)
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
