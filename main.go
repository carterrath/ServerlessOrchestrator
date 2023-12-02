package main

import (
	//"github.com/GoKubes/ServerlessOrchestrator/api"
	"github.com/GoKubes/ServerlessOrchestrator/dataAccess"
	"github.com/GoKubes/ServerlessOrchestrator/db"
)

func main() {
	// Start the API server
	//api.APIStart()

	db.CreateDatabase()
	dataAccess.ConnectToDB()
}
