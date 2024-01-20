package main

import (
	"github.com/GoKubes/ServerlessOrchestrator/DataAccess"
)

func main() {
	// Start the API server
	//api.APIStart()

	DataAccess.CreateDatabase()
	dao := &DataAccess.MicroservicesDAO{}
	dao.ConnectToDB()
}
