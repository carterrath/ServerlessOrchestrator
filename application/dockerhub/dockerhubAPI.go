package dockerhub

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateAndPushImage(backendName string) (string, error) {
	destinationPath := "/Users/carterrath/Documents/Fall2023/SE490/ServerlessOrchestrator/application/microholder/" + backendName

	dockerfile := filepath.Join(destinationPath, "Dockerfile")

	dockerRepository := "carterrath/serverless-orchestrator"
	token := os.Getenv("DOCKERHUB_TOKEN")

	// Authenticate with Docker Hub
	loginCmd := exec.Command("docker", "login", "-u", "carterrath", "--password-stdin")
	loginCmd.Stdin = strings.NewReader(token)
	loginCmd.Stdout = os.Stdout
	loginCmd.Stderr = os.Stderr
	if err := loginCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to authenticate with Docker Hub: %v", err)
	}

	// Build the Docker image
	buildCmd := exec.Command("docker", "build", "-t", backendName, "-f", dockerfile, destinationPath)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build Docker image: %v", err)
	}

	// Tag the Docker image with the custom tag
	tagCmd := exec.Command("docker", "tag", backendName, dockerRepository+":"+backendName)
	tagCmd.Stdout = os.Stdout
	tagCmd.Stderr = os.Stderr
	if err := tagCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to tag Docker image: %v", err)
	}

	// Push the built image to Dockerhub
	pushCmd := exec.Command("docker", "push", dockerRepository+":"+backendName)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to push Docker image to Dockerhub: %v", err)
	}

	// Get the digest of the pushed image
	imageInfoCmd := exec.Command("docker", "inspect", "--format='{{index .RepoDigests 0}}'", dockerRepository+":"+backendName)
	output, err := imageInfoCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get image digest: %v", err)
	}

	digest := strings.TrimSpace(strings.Trim(string(output), "'"))
	return digest, nil
}
