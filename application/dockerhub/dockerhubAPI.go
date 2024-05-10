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

	dockerRepository := os.Getenv("DOCKERHUB_REPO")

	// Build the Docker image
	buildCmd := exec.Command("docker", "build", "-t", dockerRepository+":"+backendName, "-f", dockerfile, destinationPath)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build Docker image: %v", err)
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

	// Convert output to a string and trim space for safety
	outputStr := strings.TrimSpace(string(output))

	// Remove any leading and trailing single quotes or newlines that may encapsulate the output
	digest := strings.Trim(outputStr, "'\n")

	// Remove the image tagged with the Docker repository name and backendName
	removeRepoCmd := exec.Command("docker", "rmi", dockerRepository+":"+backendName)
	removeRepoCmd.Stdout = os.Stdout
	removeRepoCmd.Stderr = os.Stderr
	if err := removeRepoCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to remove image %s: %v", dockerRepository+":"+backendName, err)
	}

	return digest, nil

}

func RunImageFromDockerHub(imageDigest, backendName string, port int) error {
	repositoryName := os.Getenv("DOCKERHUB_REPO")

	image := repositoryName + ":" + backendName

	pullCmd := exec.Command("docker", "pull", image)
	output, err := pullCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error pulling docker image: %s, output: %s", err, string(output))
	}

	// Informing the user about the pull success
	fmt.Println("Successfully pulled image:", imageDigest)

	// Run the image locally on the specified port
	// Note: Docker run command expects the port mapping in the format "hostPort:containerPort"
	runCmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%d:3000", port), "--name", backendName, image)
	runOutput, runErr := runCmd.CombinedOutput()
	if runErr != nil {
		return fmt.Errorf("error running docker image: %s, output: %s", runErr, string(runOutput))
	}

	// Informing the user that the image is running
	fmt.Println("Successfully running image:", image, "on port", port)

	return nil
}

func StopImage(backendName string) error {
	repositoryName := os.Getenv("DOCKERHUB_REPO")

	image := repositoryName + ":" + backendName
	// Stop the running container
	stopCmd := exec.Command("docker", "stop", backendName)
	stopOutput, stopErr := stopCmd.CombinedOutput()
	if stopErr != nil {
		return fmt.Errorf("error stopping container: %s, output: %s", stopErr, string(stopOutput))
	}
	fmt.Println("Successfully stopped container:", backendName)

	// Remove the stopped container
	removeCmd := exec.Command("docker", "rm", backendName)
	removeOutput, removeErr := removeCmd.CombinedOutput()
	if removeErr != nil {
		return fmt.Errorf("error removing container: %s, output: %s", removeErr, string(removeOutput))
	}
	fmt.Println("Successfully removed container:", backendName)

	// Remove the image
	removeImageCmd := exec.Command("docker", "rmi", image)
	removeImageOutput, removeImageErr := removeImageCmd.CombinedOutput()
	if removeImageErr != nil {
		return fmt.Errorf("error removing image: %s, output: %s", removeImageErr, string(removeImageOutput))
	}
	fmt.Println("Successfully removed image:", image)

	return nil
}
