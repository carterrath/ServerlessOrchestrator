package main

import (
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
)

func main() {
	// Start the API server
	//api.APIStart()

	dataaccess.CreateDatabase()
	//dao := &dataaccess.MicroservicesDAO{}
	//dao.ConnectToDB()
}
