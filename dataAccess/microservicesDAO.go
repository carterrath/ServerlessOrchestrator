package dataAccess

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	/*
		We will test this once we set up a sample GitHub! :D
		To test this we will use a pre-made repository but repoLink will be provided through the UI
		The Developer would provide their repoLink and we would clone it and run the SQL file
	*/
	repoLink := "https://github.com/JaclynW/testRepo.git"

	// Temporary local path to clone and store the repository's contents.
	clonePath := "./tempRepo"

	// Clone the provided GitHub repository.
	cloneRepo(repoLink, clonePath)

	// Assuming the repository has an SQL file with commands to insert data.
	// At this point, you would typically read this file, parse its contents,
	// and execute any relevant SQL commands against your database.

	// Before adding the microservice, check if it already exists in the database.
	if exists := checkMicroserviceExists(database, microName, repoLink); exists {
		fmt.Println("Microservice with the provided name and repository link already exists in the database.")
		return
	}

	// Insert the microservice's information into the database.
	/*
		For our testRepo, the microservice name is "services"
		We will ask user to input the microservice name which is stored in microName

		We will eventually need to figure out how to grab teh microservice name JUST from the repoLink
	*/
	addMicroservice(database, microName, repoLink)

	// Cleanup: Remove the cloned repository from local storage to free up space.
	os.RemoveAll(clonePath)

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
