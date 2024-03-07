package github

import (
	"errors"
	"os"
	"os/exec"
)

func CloneRepositoryUsingCommand(repoURL, backendName, filePath string) error {
	destinationPath := filePath + backendName
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
