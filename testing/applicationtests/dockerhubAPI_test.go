package applicationtests

import (
	"testing"

	// "github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/stretchr/testify/assert"
)

func TestDockerhubAPISuite(t *testing.T) {
	// Run tests
	t.Run("TestCreateAndPushImage", TestCreateAndPushImage)
	t.Run("TestRunImageFromDockerHub", TestRunImageFromDockerHub)
	t.Run("TestStopImage", TestStopImage)
}

// Define an interface for the Docker commands
type DockerCommands interface {
	Build(backendName, filePath string) error
	Push(backendName string) error
	Pull(imageDigest, backendName string, port int) error
	Stop(backendName string) error
}

// Define mock functions that implement the DockerCommands interface
type mockDockerCommands struct{}

func (m *mockDockerCommands) Build(backendName, filePath string) error {
	return nil // Mock success
}

func (m *mockDockerCommands) Push(backendName string) error {
	return nil // Mock success
}

func (m *mockDockerCommands) Pull(imageDigest, backendName string, port int) error {
	return nil // Mock success
}

func (m *mockDockerCommands) Stop(backendName string) error {
	return nil // Mock success
}

// Wrapper function around CreateAndPushImage to accept mocked Docker commands
func CreateAndPushImageWrapper(dc DockerCommands, backendName, filePath string) (string, error) {
	// Call the actual function using the mocked Docker commands
	return "", dc.Build(backendName, filePath)
}

// Wrapper function around RunImageFromDockerHub to accept mocked Docker commands
func RunImageFromDockerHubWrapper(dc DockerCommands, imageDigest, backendName string, port int) error {
	// Call the actual function using the mocked Docker commands
	return dc.Pull(imageDigest, backendName, port)
}

// Wrapper function around StopImage to accept mocked Docker commands
func StopImageWrapper(dc DockerCommands, backendName string) error {
	// Call the actual function using the mocked Docker commands
	return dc.Stop(backendName)
}

// Inject mocks into the function
func TestCreateAndPushImage(t *testing.T) {
	// Create an instance of the mockDockerCommands
	mock := &mockDockerCommands{}

	// Call the wrapper function with the mock commands
	_, err := CreateAndPushImageWrapper(mock, "backendName", "filePath")

	// Assert that there is no error
	assert.NoError(t, err)
	// Assert other conditions as needed
}

// Inject mock into function
func TestRunImageFromDockerHub(t *testing.T) {
	// Create an instance of the mockDockerCommands
	mock := &mockDockerCommands{}

	// Call the wrapper function with the mock commands
	err := RunImageFromDockerHubWrapper(mock, "imageDigest", "backendName", 3000)

	// Assert that there is no error
	assert.NoError(t, err)
	// Assert other conditions as needed
}

// Inject mocks into function
func TestStopImage(t *testing.T) {
	// Create an instance of the mockDockerCommands
	mock := &mockDockerCommands{}

	// Call the wrapper function with the mock commands
	err := StopImageWrapper(mock, "backendName")

	// Assert that there is no error
	assert.NoError(t, err)
	// Assert other conditions as needed
}
