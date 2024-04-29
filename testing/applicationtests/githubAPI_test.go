package applicationtests

import (
	"log"
	"os"
	"os/exec"

	"path/filepath"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/GoKubes/ServerlessOrchestrator/application/github"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Set testPath to be relative to the current working directory
	testPath = filepath.Join(cwd, "/application/microholder/")
}

var (
	testURL  = "https://github.com/ruthijimenez/shopping-cart.git"
	testPath string
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
	//os.RemoveAll(testPath + "testBackendName")
	//os.RemoveAll(testPath + "testBackendName2")

	// delete the /application/microholder directory
	err := os.RemoveAll("/application/")
	if err != nil {
		log.Printf("Error removing directory: %v", err)
	}
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
	teardown()
}

func TestCloneRepositoryUsingCommandPass(t *testing.T) {
	assert.NoError(t, github.CloneRepositoryUsingCommand(testURL, "testBackendName", testPath))

	os.RemoveAll(testPath + "testBackendName")
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
	teardown()
}

func TestGetLastestPushDatePass(t *testing.T) {
	date, err := github.GetLatestPushDate(testURL, "testBackendName2", testPath)
	assert.NoError(t, err)
	assert.NotNil(t, date)

	os.RemoveAll(testPath + "testBackendName2")
}
