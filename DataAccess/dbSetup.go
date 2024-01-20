package DataAccess

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	database, err := sql.Open("sqlite3", "./microservice.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	/*
	 Create a table for microservices
	 Each microservice will have a service hook, build script, and placeholder
	*/
	createMicroservicesTable := `CREATE TABLE IF NOT EXISTS microservices (
		"id" INTEGER PRIMARY KEY, 
		"name" TEXT UNIQUE,
		"service_hook" TEXT, 
		"build_script" TEXT,
		"place_holder" TEXT
	);`

	_, err = database.Exec(createMicroservicesTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
