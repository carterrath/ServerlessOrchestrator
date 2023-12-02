package dataAccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() {
	database, err := sql.Open("sqlite3", "./microservice.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	var microName string
	fmt.Print("Enter your Microservice name: ")
	if _, err := fmt.Scanln(&microName); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Retrieve the repository link associated with the microservice from the database.
	repoLink := getRepoLinkFromDatabase(database, microName)

	// Check if the retrieved repository link is empty or not found.
	if repoLink == "" {
		fmt.Println("Microservice not found in the database or repository link is empty.")
		// Allow the user to upload a new microservice and repository link.
		uploadNewMicroservice(database, microName)
		return
	}

	// Temporary local path to clone and store the repository's contents.
	clonePath := "./tempRepo"

	// Check if the repository already exists in the specified path.
	if _, err := os.Stat(clonePath); err == nil {
		// The repository already exists, so update it.
		updateRepository(clonePath)
	} else {
		// Clone the repository if it doesn't exist.
		cloneRepo(repoLink, clonePath)
	}

	// Assuming the repository has an SQL file with commands to insert data.
	// At this point, you would typically read this file, parse its contents,
	// and execute any relevant SQL commands against your database.

	// Before adding the microservice, check if it already exists in the database.
	if exists := checkMicroserviceExists(database, microName, repoLink); exists {
		fmt.Println("Microservice with the provided name and repository link already exists in the database.")

		// Ask the user whether to run the existing microservice.
		var runMicroservice string
		fmt.Print("Do you want to run this microservice? (yes/no): ")
		if _, err := fmt.Scanln(&runMicroservice); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		if strings.ToLower(runMicroservice) == "yes" {
			// Run the existing microservice.
			fmt.Printf("Running existing microservice: %s\n", microName)
			// Discover and execute scripts from the cloned repository
			discoverAndExecuteScripts(clonePath)
			// Cleanup: Remove the cloned repository from local storage to free up space.
			os.RemoveAll(clonePath)
		} else {
			// User chose not to run the microservice, so terminate.
			fmt.Println("Microservice not executed. Terminating.")
			return
		}
		return
	} else {
		cloneRepo(repoLink, clonePath)
		addMicroservice(database, microName, repoLink)

		// Discover and execute scripts from the cloned repository
		discoverAndExecuteScripts(clonePath)

		// Cleanup: Remove the cloned repository from local storage to free up space.
		os.RemoveAll(clonePath)
	}

}

// cloneRepo function uses the go-git package to clone a given GitHub repository
// into a specified local path.
func cloneRepo(url, path string) {
	// The PlainClone function clones a repository into the path.
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,       // Repository URL
		Progress: os.Stdout, // Display progress on standard output
	})
	if err != nil {
		log.Fatalf("Error cloning the repository: %v", err)
	}
}

// addMicroservice function inserts the microservice's name and its associated
// GitHub repository link into a database table.
func addMicroservice(database *sql.DB, name, repoLink string) {
	// Prepare an SQL statement for insertion.
	statement, _ := database.Prepare("INSERT INTO microservices (build_script, placeholder) VALUES (?, ?)")
	statement.Exec(name, repoLink) // Execute the statement with provided values.
}

// checkMicroserviceExists function checks if a microservice with the given name
// and repository link already exists in the database.
func checkMicroserviceExists(database *sql.DB, name, repoLink string) bool {
	// Prepare a SELECT query to find the microservice with the given name and repo link.
	row := database.QueryRow("SELECT count(*) FROM microservices WHERE build_script = ? AND placeholder = ?", name, repoLink)

	var count int
	if err := row.Scan(&count); err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return false
	}

	// Return true if count is more than 0, indicating the microservice exists.
	return count > 0
}

// uploadNewMicroservice allows the user to upload a new microservice and its repository link.
func uploadNewMicroservice(database *sql.DB, microName string) {
	var newRepoLink string
	fmt.Printf("Enter the repository link for microservice '%s': ", microName)
	if _, err := fmt.Scanln(&newRepoLink); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Insert the new microservice's information into the database.
	addMicroservice(database, microName, newRepoLink)
	fmt.Println("New microservice added to the database.")
}

// getRepoLinkFromDatabase retrieves the repository link associated with a microservice from the database.
func getRepoLinkFromDatabase(database *sql.DB, microName string) string {
	var repoLink string
	row := database.QueryRow("SELECT placeholder FROM microservices WHERE build_script = ?", microName)
	if err := row.Scan(&repoLink); err != nil {
		// Handle the error or return an empty string if not found.
		fmt.Printf("Error retrieving repository link: %v\n", err)
		return ""
	}
	return repoLink
}

func updateRepository(repoPath string) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Error opening the existing repository: %v", err)
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Error accessing the worktree: %v", err)
		return
	}

	// Pull the latest changes from the remote repository
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("Error pulling changes: %v", err)
		return
	}

	fmt.Println("Repository updated successfully.")
}

func discoverAndExecuteScripts(repoPath string) {
	dirEntries, err := os.ReadDir(repoPath)
	if err != nil {
		log.Fatalf("Error reading repository directory: %v", err)
		return
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue // Skip directories
		}

		if isScript(dirEntry.Name()) {
			fmt.Printf("Executing script: %s\n", dirEntry.Name())
			scriptPath := filepath.Join(repoPath, dirEntry.Name())
			cmd := exec.Command("bash", "-c", scriptPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				log.Fatalf("Error executing script %s: %v", dirEntry.Name(), err)
			}
		}
	}
}

func isScript(filename string) bool {
	// Any file with a .sh extension is a script.
	return strings.HasSuffix(filename, ".sh")
}
