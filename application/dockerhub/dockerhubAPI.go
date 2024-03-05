package dockerhub

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

// CreateAndPushImage builds a Docker image from the specified directory and Dockerfile,
// tags it with the given image name, and pushes it to Dockerhub using the provided access token.
func CreateAndPushImageDockerClient(backendName, dockerfile string) error {
	destinationPath := "application/microholder/" + backendName
	dockerRepository := "carterrath/serverless-orchestrator:latest"
	token := os.Getenv("DOCKERHUB_TOKEN")

	// Initialize Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return fmt.Errorf("failed to initialize Docker client: %v", err)
	}

	// Build the Docker image
	imageBuildResponse, err := dockerClient.ImageBuild(context.Background(),
		getBuildContext(destinationPath), types.ImageBuildOptions{
			Tags:       []string{dockerRepository},
			Dockerfile: dockerfile,
		})
	if err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}
	defer imageBuildResponse.Body.Close()
	io.Copy(os.Stdout, imageBuildResponse.Body)

	// Push the built image to Dockerhub
	pushResponse, err := dockerClient.ImagePush(context.Background(), dockerRepository,
		types.ImagePushOptions{
			RegistryAuth: "Bearer " + token,
		})
	if err != nil {
		return fmt.Errorf("failed to push Docker image to Dockerhub: %v", err)
	}
	defer pushResponse.Close()
	io.Copy(os.Stdout, pushResponse)

	return nil
}

// Function to prepare the build context
func getBuildContext(contextDir string) io.Reader {
	tarBuildContext, err := archiveDirectory(contextDir)
	if err != nil {
		panic(err)
	}
	return tarBuildContext
}

// Function to archive the directory for building
func archiveDirectory(dirPath string) (io.Reader, error) {
	tarOptions := &archive.TarOptions{}
	return archive.TarWithOptions(dirPath, tarOptions)
}

func CreateAndPushImage(backendName string) error {
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
		return fmt.Errorf("failed to authenticate with Docker Hub: %v", err)
	}

	// Build the Docker image
	buildCmd := exec.Command("docker", "build", "-t", backendName, "-f", dockerfile, destinationPath)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build Docker image: %v", err)
	}

	// Tag the Docker image with the custom tag
	// Tag the Docker image built from backendName with the custom tag
	tagCmd := exec.Command("docker", "tag", backendName, dockerRepository+":"+backendName)

	tagCmd.Stdout = os.Stdout
	tagCmd.Stderr = os.Stderr
	if err := tagCmd.Run(); err != nil {
		return fmt.Errorf("failed to tag Docker image: %v", err)
	}

	// Push the built image to Dockerhub
	pushCmd := exec.Command("docker", "push", dockerRepository+":"+backendName)
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("failed to push Docker image to Dockerhub: %v", err)
	}

	return nil
}
