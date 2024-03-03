package dataaccess

import (
	"errors"
	"os"
	"os/exec"

	git "gopkg.in/src-d/go-git.v4"
)

func CloneRepositoryUsingCommand(repoURL, backendName string) error {
	destinationPath := "application/microholder/" + backendName
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

func CloneRepositoryUsingGit(url string) error {
	destinationPath := "application/microholder"
	if _, err := os.Stat(destinationPath); err == nil {
		return errors.New("destination directory already exists")
	}
	// The PlainClone function clones a repository into the path.
	_, err := git.PlainClone(destinationPath, false, &git.CloneOptions{
		URL:      url,       // Repository URL
		Progress: os.Stdout, // Display progress on standard output
	})
	if err != nil {
		return err
	}
	return nil
}
