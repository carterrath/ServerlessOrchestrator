package dataaccess

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/GoKubes/ServerlessOrchestrator/business"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres_SDK struct {
	db  *gorm.DB
	dsn string
}

func CreateDatabase() {
	// PostgreSQL connection string
	// "host=localhost user=username password=password dbname=database_name sslmode=disable"
	// Adjust with your actual credentials and database name
	// store credentials in an environment variable
	Username := os.Getenv("POSTGRES_USERNAME")
	//Password := os.Getenv("POSTGRES_PASSWORD")
	Password := ""
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DB := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DB, Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// pooling?
	if dbC, err := db.DB(); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	} else {
		dbC.SetMaxIdleConns(22)
		dbC.SetMaxOpenConns(22)
		dbC.SetConnMaxLifetime(time.Hour)
	}

	// AutoMigrate will create or migrate your tables according to the struct
	err = db.AutoMigrate(&business.Microservice{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

}
