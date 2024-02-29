package application

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func IsPublicRepo(repoLink string) bool {
	token := os.Getenv("GITHUB_TOKEN")

	// Extract owner and repository name from the URL
	parts := strings.Split(repoLink, "/")
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
		log.Fatalf("Error getting repository: %v", err)
	}
	if *repo.Private {
		return false
	}

	return true
}
