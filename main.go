package main

import (
	"fmt"
	"log"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func main() {

	// dataaccess.CreateDatabase()
	db := dataaccess.CreateDatabase()

	// Create an instance of MicroservicesDAO
	microservicesDAOpq := dataaccess.NewMicroservicesDAO(db)

	// Example: Insert a record into the Microservice table
	microservice := business.Microservice{
		Name:        "catalog",
		ServiceHook: "http://my-service-num2.com",
		BuildScript: "npm run build",
		PlaceHolder: "placeholder",
	}

	// Create record without DAO
	// result := db.Create(&microservice)
	// fmt.Println(result.RowsAffected) //For console, to check if record was created
	// if result.Error != nil {
	// 	log.Fatalf("failed to insert record: %v", result.Error)
	// }

	// Use the DAO to insert the microservice record
	err := microservicesDAOpq.Insert(microservice)
	if err != nil {
		log.Fatalf("failed to insert microservice: %v", err)
	} else {
		fmt.Println("Microservice inserted successfully")
	}
}
