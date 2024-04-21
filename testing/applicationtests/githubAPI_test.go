package applicationtests

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/stretchr/testify/assert"
)

var (
	testURL  = "https://github.com/ruthijimenez/shopping-cart.git"
	testPath = "/Users/ruthjimenez/Documents/GitHub/ServerlessOrchestrator/application/microholder/"
)

func TestGithubAPISuite(t *testing.T) {
	t.Run("CloneRepositoryUsingCommand", TestCloneRepositoryUsingCommand)
	t.Run("CloneRepositoryUsingCommandPass", TestCloneRepositoryUsingCommandPass)
	t.Run("GetLatestPushDate", TestGetLatestPushDate)
	t.Run("GetLastestPushDatePass", TestGetLastestPushDatePass)

	// teardown
	teardown()
}

func teardown() {
	// delete the cloned repo
	os.RemoveAll(testPath + "testBackendName")
	os.RemoveAll(testPath + "testBackendName2")
}

func TestCloneRepositoryUsingCommand(t *testing.T) {
	// Mock os.Stat
	monkey.Patch(os.Stat, func(string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	})
	defer monkey.Unpatch(os.Stat)

	// Mock exec.Command
	monkey.Patch(exec.Command, func(name string, arg ...string) *exec.Cmd {
		return &exec.Cmd{}
	})
	defer monkey.Unpatch(exec.Command)

	assert.Error(t, github.CloneRepositoryUsingCommand("repoURL", "backendName", "filePath"))

}

func TestCloneRepositoryUsingCommandPass(t *testing.T) {
	assert.NoError(t, github.CloneRepositoryUsingCommand(testURL, "testBackendName", testPath))
}

func TestGetLatestPushDate(t *testing.T) {
	// Mock os.Stat
	monkey.Patch(os.Stat, func(string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	})
	defer monkey.Unpatch(os.Stat)

	// Mock exec.Command
	monkey.Patch(exec.Command, func(name string, arg ...string) *exec.Cmd {
		cmd := &exec.Cmd{}
		monkey.PatchInstanceMethod(reflect.TypeOf(cmd), "Output", func(*exec.Cmd) ([]byte, error) {
			return []byte("2022-01-01"), nil
		})
		return cmd
	})
	defer monkey.Unpatch(exec.Command)

	// Call GetLatestPushDate and separate the return values
	_, err := github.GetLatestPushDate("repoURL", "backendName", "filePath")

	// Check that error was returned
	assert.Error(t, err)
}

func TestGetLastestPushDatePass(t *testing.T) {
	date, err := github.GetLatestPushDate(testURL, "testBackendName2", testPath)
	assert.NoError(t, err)
	assert.NotNil(t, date)
}
