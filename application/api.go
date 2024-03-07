/*
Package application provides the API server for the application.

The API server is a RESTful API that allows clients to interact with the application.

The API server is built using the Gin web framework.

The API server provides the following endpoints:

- GET /items: Get all items
- POST /items: Add a new item

API provides the following Devloper endpoints:
- POST /submit-repo: Submit a GitHub repository for analysis
  - Check public or private
  - If public, check for duplicates in database
  - If private, return error
  - If no duplicates, add to database
  - If duplicates, return error

- GET /microservices: Get all microservices
  - Return all microservices in the database
  - Return error if database is empty

- GET /microservices/:id: Get a microservice by ID

The API server is started by calling the APIStart function.
*/
package application

import (
	"net/http"

	"fmt"

	"github.com/GoKubes/ServerlessOrchestrator/application/routes/microservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/routes/runmicroservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/routes/users"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

func Init(dao *dataaccess.MicroservicesDAO, userdao *dataaccess.UserDAO) error {
	router := gin.Default()

	// Add CORS middleware to allow requests from http://localhost:5173
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	handleRoutes(router, dao, userdao)

	err := router.Run("localhost:8080")
	if err != nil {
		return err
	}

	fmt.Println("Server is now listening on localhost:8080")

	return nil
}

func handleRoutes(router *gin.Engine, dao *dataaccess.MicroservicesDAO, userdao *dataaccess.UserDAO) {
	//MicroserviceRouter := router.Group("/microservice")
	router.GET("/microservice", func(c *gin.Context) {
		microservice.GetAllMicroservices(c, dao, userdao)
	})
	router.POST("/microservice", func(c *gin.Context) {
		microservice.UploadMicroservice(c, dao)
	})
	router.POST("/signup/developer", func(c *gin.Context) {
		users.CreateDeveloper(c, userdao)
	})
	router.POST("/signup/consumer", func(c *gin.Context) {
		users.CreateConsumer(c, userdao)
	})
	router.POST("/login/developer", func(c *gin.Context) {
		users.Login(c, userdao)
	})
	router.POST("/login/consumer", func(c *gin.Context) {
		users.Login(c, userdao)
	})
	router.POST("/runmicroservice", func(c *gin.Context) {
		runmicroservice.RunMicroservice(c, dao)
	})
}
