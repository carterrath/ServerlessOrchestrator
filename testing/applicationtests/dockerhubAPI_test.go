package applicationtests

import (
	"os/exec"
	"testing"

	"bou.ke/monkey"
	"github.com/GoKubes/ServerlessOrchestrator/application/dockerhub"
	"github.com/stretchr/testify/assert"
)

func TestDockerhubAPISuite(t *testing.T) {
	// Run tests
	t.Run("TestCreateAndPushImage", TestCreateAndPushImage)
	t.Run("TestRunImage", TestRunImage)
	t.Run("TestStopImage", TestStopImage)
}

func TestCreateAndPushImage(t *testing.T) {
	// Patch exec.Command
	monkey.Patch(exec.Command, func(name string, arg ...string) *exec.Cmd {
		// Check the command and args if needed, return appropriate mock command
		if name == "docker" && arg[0] == "build" {
			// Mock behavior for docker build command
			return &exec.Cmd{}
		} else if name == "docker" && arg[0] == "push" {
			// Mock behavior for docker push command
			return &exec.Cmd{}
		} else {
			// Return nil for other commands
			return nil
		}
	})
	defer monkey.Unpatch(exec.Command) // Unpatch after the test

	// Call the function under test
	_, err := dockerhub.CreateAndPushImage("backendName", "filePath")

	// Assert that there is an error
	assert.Error(t, err)
}

func TestRunImage(t *testing.T) {
	// Patch exec.Command
	monkey.Patch(exec.Command, func(name string, arg ...string) *exec.Cmd {
		// Check the command and args if needed, return appropriate mock command
		if name == "docker" && arg[0] == "pull" {
			// Mock behavior for docker pull command
			return &exec.Cmd{}
		} else if name == "docker" && arg[0] == "run" {
			// Mock behavior for docker run command
			return &exec.Cmd{}
		} else {
			// Return nil for other commands
			return nil
		}
	})
	defer monkey.Unpatch(exec.Command) // Unpatch after the test

	// Call the function under test
	err := dockerhub.RunImageFromDockerHub("imageDigest", "backendName", 8080)

	// Assert that there is an error
	assert.Error(t, err)
}

func TestStopImage(t *testing.T) {
	// Patch exec.Command
	monkey.Patch(exec.Command, func(name string, arg ...string) *exec.Cmd {
		// Check the command and args if needed, return appropriate mock command
		if name == "docker" && arg[0] == "stop" {
			// Mock behavior for docker stop command
			return &exec.Cmd{}
		} else if name == "docker" && (arg[0] == "rm" || arg[0] == "rmi") {
			// Mock behavior for docker rm and rmi commands
			return &exec.Cmd{}
		} else {
			// Return nil for other commands
			return nil
		}
	})
	defer monkey.Unpatch(exec.Command) // Unpatch after the test

	// Call the function under test
	err := dockerhub.StopImage("backendName")

	// Assert that there is an error
	assert.Error(t, err)
}
