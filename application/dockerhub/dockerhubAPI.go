package dockerhub

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateAndPushImage(backendName, filePath string) (string, error) {
	destinationPath := filePath + backendName

	dockerfile := filepath.Join(destinationPath, "Dockerfile")

	dockerRepository := "carterrath/serverless-orchestrator"
	// token := os.Getenv("DOCKERHUB_TOKEN")

	// Authenticate with Docker Hub
	// loginCmd := exec.Command("docker", "login", "-u", "jaclynw", "--password-stdin")
	// loginCmd.Stdin = strings.NewReader(token)
	// loginCmd.Stdout = os.Stdout
	// loginCmd.Stderr = os.Stderr
	// if err := loginCmd.Run(); err != nil {
	// 	return "", fmt.Errorf("failed to authenticate with Docker Hub: %v", err)
	// }

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

	// Remove the local image tagged with backendName
	removeLocalCmd := exec.Command("docker", "rmi", backendName)
	removeLocalCmd.Stdout = os.Stdout
	removeLocalCmd.Stderr = os.Stderr
	if err := removeLocalCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to remove local image %s: %v", backendName, err)
	}

	// Remove the image tagged with the Docker repository name and backendName
	removeRepoCmd := exec.Command("docker", "rmi", dockerRepository+":"+backendName)
	removeRepoCmd.Stdout = os.Stdout
	removeRepoCmd.Stderr = os.Stderr
	if err := removeRepoCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to remove image %s: %v", dockerRepository+":"+backendName, err)
	}

	return digest, nil

}

func RunImageFromDockerHub(imageDigest string, port int) error {
	repositoryName := "carterrath/serverless-orchestrator"
	// Use Docker CLI to inspect the manifest of the image on Docker Hub
	inspectCmd := exec.Command("docker", "manifest", "inspect", repositoryName)
	output, err := inspectCmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Convert the output to string
	manifestJSON := string(output)

	// Check if the image digest is found in the search results
	if !strings.Contains(manifestJSON, imageDigest) {
		return fmt.Errorf("image with digest %s not found in Docker Hub repository %s", imageDigest, repositoryName)
	}

	// Pull the image from Docker Hub
	pullCmd := exec.Command("docker", "pull", repositoryName)
	if err := pullCmd.Run(); err != nil {
		return err
	}

	// Run the image locally on the specified port
	runCmd := exec.Command("docker", "run", "-p", fmt.Sprintf("%d:8080", port), repositoryName)
	if err := runCmd.Run(); err != nil {
		return err
	}

	return nil
}
