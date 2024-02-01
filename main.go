package main

import (
	// "github.com/GoKubes/ServerlessOrchestrator/application"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func main() {
	// Start the API server
	// application.APIStart()

	dataaccess.CreateDatabase()
	//dao := &dataaccess.MicroservicesDAO{}
	//dao.ConnectToDB()
}
