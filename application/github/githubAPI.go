package github

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func CloneRepositoryUsingCommand(repoURL, backendName string) error {
	destinationPath := "/Users/jwalsh/Dev/CSUSM/SE490 Capstone/ServerlessOrchestrator/application/microholder/" + backendName
	// Check if the destination directory already exists

	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("destination directory already exists")
	}

	// Run the git clone command
	cmd := exec.Command("git", "clone", repoURL, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command and check for errors
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// getLatestPushDate fetches the latest push date of a GitHub repo.
func GetLatestPushDate(repoURL, backendName string) (string, error) {
	destinationPath := "/Users/jwalsh/Dev/CSUSM/SE490 Capstone/ServerlessOrchestrator/application/microholder/" + backendName

	if err := CloneRepositoryUsingCommand(repoURL, backendName); err != nil {
		return "", err
	}
	println("Error: Up here!! ")
	// Get the latest commit date.
	cmd := exec.Command("git", "-C", destinationPath, "log", "-1", "--format=%cd")
	output, err := cmd.Output()
	if err != nil {
		println("Error: I'm here!! ", err)
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
