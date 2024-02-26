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

// PostgresUserDB is a struct that holds the GORM database connection and the connection string for the user database.
type PostgresUserDB struct {
	db  *gorm.DB
	dsn string
}

// CreateUserDatabase initializes the PostgreSQL user database connection and performs necessary setup.
func CreateUserDatabase() *gorm.DB {
	// PostgreSQL connection string for the user database
	Username := os.Getenv("POSTGRES_USERNAME")
	// Password := os.Getenv("POSTGRES_PASSWORD")
	Password := ""
	Host := os.Getenv("POSTGRES_HOST")
	Port := os.Getenv("POSTGRES_PORT")
	DBName := os.Getenv("POSTGRES_USER_DB")

	// Construct the data source name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", Host, Username, Password, DBName, Port)

	// Open a GORM database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to user database: %v", err)
	}

	// Connection pooling configuration
	if dbC, err := db.DB(); err != nil {
		log.Fatalf("failed to connect to user database: %v", err)
	} else {
		dbC.SetMaxIdleConns(22)
		dbC.SetMaxOpenConns(22)
		dbC.SetConnMaxLifetime(time.Hour)
	}

	// AutoMigrate will create or migrate your tables according to the struct
	err = db.AutoMigrate(&business.User{})
	if err != nil {
		log.Fatalf("failed to migrate user database: %v", err)
	}

	return db
}
