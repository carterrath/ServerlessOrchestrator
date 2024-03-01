package dataaccess

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func IsPublicRepo(repoLink string) (bool, error) {
	token := os.Getenv("GITHUB_TOKEN")

	// Verify the URL format
	if !strings.HasPrefix(repoLink, "https://github.com/") {
		return false, fmt.Errorf("Invalid GitHub repository URL")
	}

	// Extract owner and repository name from the URL
	parts := strings.Split(repoLink, "/")
	if len(parts) < 3 {
		return false, fmt.Errorf("Invalid GitHub repository URL")
	}
	owner := parts[len(parts)-2]
	repoName := parts[len(parts)-1]

	// Create a new GitHub client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Check if the repository is public or private
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		// Print out the actual response body when an error occurs
		if errResp, ok := err.(*github.ErrorResponse); ok {
			return false, fmt.Errorf("error getting repository: %d - %s", errResp.Response.StatusCode, errResp.Message)
		}
		return false, fmt.Errorf("error getting repository: %v", err)
	}

	if repo == nil {
		return false, fmt.Errorf("repository not found")
	}

	if *repo.Private {
		return false, nil
	}

	return true, nil
}
