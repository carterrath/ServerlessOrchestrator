package main

import (
	"context"
	"fmt"
	"log"

	"github.com/GoKubes/ServerlessOrchestrator/application"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/joho/godotenv"
)

// Define a struct to hold dependencies

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	db := dataaccess.CreateDatabase()
	dao := dataaccess.NewMicroservicesDAO(db)
	userdao := dataaccess.NewUserDAO(db)

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Create an ECS client
	ecsClient := ecs.NewFromConfig(cfg)

	// Initialize the Route 53 client
	r53Client := route53.NewFromConfig(cfg)

	if err := application.Init(dao, userdao, ecsClient, r53Client); err != nil {
		panic(err)
	}
}
