package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func main() {

	// dataaccess.CreateDatabase()
	db := dataaccess.CreateDatabase()

	// Create an instance of MicroservicesDAO
	microservicesDAOpq := dataaccess.NewMicroservicesDAO(db)

	// Example: Insert a record into the Microservice table
	// microservice := business.Microservice{
	// 	Name:        "catalog",
	// 	ServiceHook: "http://my-service-num2.com",
	// 	BuildScript: "npm run build",
	// 	PlaceHolder: "placeholder",
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
	// 	fmt.Printf("ID: %d, Name: %s, ServiceHook: %s, BuildScript: %s, Placeholder: %s\n", ms.ID, ms.Name, ms.ServiceHook, ms.BuildScript, ms.PlaceHolder)
	// }

	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Submit a new microservice")
		fmt.Println("2. View Microservices")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		choice := readInput()

		switch choice {
		case "1":
			submitNewMicroservice(microservicesDAOpq)
		case "2":
			viewAndExecuteMicroservices(microservicesDAOpq)
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
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
	repoLink := readInput()

	newMicroservice := business.Microservice{
		Name:        name,
		ServiceHook: repoLink,
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
		// Here, you would add your logic to execute the microservice
	} else {
		fmt.Println("Returning to main menu...")
	}
}
