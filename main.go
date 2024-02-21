package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/GoKubes/ServerlessOrchestrator/application"
	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/go-git/go-git/v5"
)

// Define a struct to hold dependencies

func main() {
	db := dataaccess.CreateDatabase()
	dao := dataaccess.NewMicroservicesDAO(db)
	// dataaccess.CreateDatabase()
	// db := dataaccess.CreateDatabase()
	if err := application.Init(dao); err != nil {
		panic(err)
	}
	// Create an instance of MicroservicesDAO
	//  microservicesDAOpq := dataaccess.NewMicroservicesDAO(db)

	// Example: Insert a record into the Microservice table
	// microservice := business.Microservice{
	// 	Name:        "catalog",
	// 	ServiceHook: "http://my-service-num2.com",
	// 	BuildScript: "npm run build",
	// 	RepoLink: "placeholder",
	// }

	// Use the DAO to insert the microservice record
	// err := microservicesDAOpq.Insert(microservice)
	// if err != nil {
	// 	log.Fatalf("failed to insert microservice: %v", err)
	// } else {
	// 	fmt.Println("Microservice inserted successfully")
	// }

	// Example: Get all records from the Microservice table
	// microservices, err := microservicesDAOpq.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get microservices: %v", err)
	// }
	// fmt.Println("Microservices retrieved successfully:")
	// for _, ms := range microservices {
	// 	fmt.Printf("ID: %d, Name: %s, ServiceHook: %s, BuildScript: %s, Placeholder: %s\n", ms.ID, ms.Name, ms.ServiceHook, ms.BuildScript, ms.RepoLink)
	// }

	// for {
	// 	fmt.Println("\nMain Menu:")
	// 	fmt.Println("1. Submit a new microservice")
	// 	fmt.Println("2. View Microservices")
	// 	fmt.Println("3. Exit")
	// 	fmt.Print("Enter your choice: ")

	// 	choice := readInput()

	// 	switch choice {
	// 	case "1":
	// 		submitNewMicroservice(microservicesDAOpq)
	// 	case "2":
	// 		viewAndExecuteMicroservices(microservicesDAOpq)
	// 	case "3":
	// 		fmt.Println("Exiting...")
	// 		return
	// 	default:
	// 		fmt.Println("Invalid choice, please try again.")
	// 	}
	// }

	//dao := &dataaccess.MicroservicesDAO{}
	//dao.ConnectToDB()
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func submitNewMicroservice(dao *dataaccess.MicroservicesDAOpq) {
	fmt.Print("Enter the name for the new microservice: ")
	name := readInput()
	fmt.Print("Enter the repository link for the new microservice: ")

	newMicroservice := business.Microservice{
		Name: name,
		// Populate other necessary fields as needed
	}

	if err := dao.Insert(newMicroservice); err != nil {
		log.Printf("Failed to add new microservice: %v", err)
	} else {
		fmt.Println("New microservice added successfully.")
	}
}

func viewAndExecuteMicroservices(dao *dataaccess.MicroservicesDAOpq) {
	microservices, err := dao.GetAll()
	if err != nil {
		log.Printf("Failed to get microservices: %v", err)
		return
	}

	if len(microservices) == 0 {
		fmt.Println("No microservices found.")
		return
	}

	fmt.Println("Microservices:")
	for _, m := range microservices {
		fmt.Printf("- %s\n", m.Name)
	}

	fmt.Println("Would you like to execute any of the microservices listed? (yes/no)")
	if strings.ToLower(readInput()) == "yes" {
		fmt.Print("Enter the name of the service to execute: ")
		serviceName := readInput()

		// Simulate executing the microservice
		fmt.Printf("Executing %s...\n", serviceName)

		// Find the microservice by name to get the repo link
		microservice, err := dao.GetByName(serviceName)
		if err != nil {
			log.Printf("Failed to find microservice '%s': %v", serviceName, err)
			return
		}

		// ServiceHook is the repoLink
		repoLink := microservice.RepoLink
		fmt.Printf("Executing %s...\n", serviceName)

		// Temporary local path to clone and store the repository's contents.
		clonePath := fmt.Sprintf("./tempRepo_%s", serviceName)

		// Check if the repository already exists in the specified path.
		if _, err := os.Stat(clonePath); os.IsNotExist(err) {
			// Clone the repository if it doesn't exist.
			cloneRepo(repoLink, clonePath)
		} else {
			// The repository already exists, so update it.
			updateRepository(clonePath)
		}

		// Execute the microservice
		fmt.Printf("Running existing microservice: %s\n", serviceName)
		// Discover and execute scripts from the cloned repository
		discoverAndExecuteScripts(clonePath)
		// Cleanup: Remove the cloned repository from local storage to free up space.
		os.RemoveAll(clonePath)

	} else {
		fmt.Println("Returning to main menu...")
	}

}

// cloneRepo function uses the go-git package to clone a given GitHub repository
// This should be in a different class
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

// Cleanup: Remove the cloned repository from local storage to free up space.
// os.RemoveAll(clonePath)
