package dataaccess

import (
	"errors"
	"os"
	"os/exec"
)

func CloneRepository(repoURL string) error {
	destinationPath := "application/microholder"
	// Check if the destination directory already exists
	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("destination directory already exists")
	}

	// Run the git clone command
	cmd := exec.Command("git", "clone", repoURL, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
