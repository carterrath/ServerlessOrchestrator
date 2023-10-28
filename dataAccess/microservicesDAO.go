package dataAccess

import (
	"database/sql"
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
	// We will test this once we set up a sample GitHub! :D
	repoLink := "https://github.com/JaclynW/testRepo.git" // Example GitHub repository link.
	// Temporary local path to clone and store the repository's contents.
	clonePath := "./tempRepo"

	// Clone the provided GitHub repository.
	cloneRepo(repoLink, clonePath)

	// Assuming the repository has an SQL file with commands to insert data.
	// At this point, you would typically read this file, parse its contents,
	// and execute any relevant SQL commands against your database.

	// ... Code to process data and insert into the database ...

	// Insert the microservice's information into the database.
	addMicroservice(database, "MicroserviceName", repoLink) // MicroserviceName changed to service hook or key?

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
