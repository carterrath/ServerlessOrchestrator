package servicestests

import (
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/application/services"
	"github.com/stretchr/testify/assert"
)

func TestUploadServiceSuite(t *testing.T) {

	// Run tests
	t.Run("TestUploadService_ValidateGithubURL", TestUploadService_ValidateGithubURL)
	t.Run("TestUploadService_GenerateBackendName", TestUploadService_GenerateBackendName)
	t.Run("TestUploadService_GetImageDigest", TestUploadService_GetImageDigest)
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
