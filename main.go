package main

import (
	"github.com/GoKubes/ServerlessOrchestrator/dataAccess"
	"github.com/GoKubes/ServerlessOrchestrator/db"
)

func main() {
	db.CreateDatabase()
	dataAccess.ConnectToDB()
}
