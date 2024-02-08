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
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/gin-gonic/gin"
)

// APIStart initializes and starts the API server
func APIStart() {
	router := gin.Default()

	// Register API routes
	RegisterRoutes(router)

	// Run the server on port 8080
	router.Run(":8080") //TODO: Change to 3000? or 3001?

}

// RegisterRoutes registers API endpoints with the provided Gin router
func RegisterRoutes(router *gin.Engine) {
	router.GET("/items", GetItems)
	router.POST("/items", AddItem)
	router.GET("/check-repo-visibility", checkRepoVisibility)
}

// ///// Test Data //////////////////
func Init(dao *dataaccess.MicroservicesDAOpq) error {
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

	handleRoutes(router, dao)

	err := router.Run("localhost:8080")
	if err != nil {
		return err
	}

	fmt.Println("Server is now listening on localhost:8080")

	return nil
}

func handleRoutes(router *gin.Engine, dao *dataaccess.MicroservicesDAOpq) {
	//MicroserviceRouter := router.Group("/microservice")
	router.GET("/microservice", func(c *gin.Context) {
		microservice.GetAllMicroservices(c, dao)
	})
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 1, Name: "Item 1"},
	{ID: 2, Name: "Item 2"},
}

// getAlbums responds with the list of all albums as JSON.

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming you generate a unique ID for the new item
	newItem.ID = len(items) + 1

	// Add the new item to the list
	items = append(items, newItem)

	c.JSON(http.StatusCreated, newItem)
}

////////////////////////////////////

func checkRepoVisibility(c *gin.Context) {
	repoURL := c.Query("repo_url")

	// Make a GET request to the GitHub API
	resp, err := http.Get(repoURL)
	if err != nil {
		// Handle error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check repository"})
		return
	}
	defer resp.Body.Close()

	// Determine visibility based on status code
	isPublic := resp.StatusCode == 200

	c.JSON(http.StatusOK, gin.H{
		"is_public": isPublic,
	})
}

//Second version of checkRepoVisibility function.
// func checkRepoVisibility(c *gin.Context) {
//     repoURL := c.Query("repo_url")

//     // Extract the owner and repo from repoURL
//     // Assuming repoURL is in the format "https://github.com/{owner}/{repo}"
//     // Extract {owner} and {repo} and use them in the GitHub API URL

//     // Use the GitHub API endpoint
//     githubAPIURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

//     // Make the request
//     resp, err := http.Get(githubAPIURL)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check repository"})
//         return
//     }
//     defer resp.Body.Close()

//     // Determine visibility based on status code
//     switch resp.StatusCode {
//     case http.StatusOK:
//         c.JSON(http.StatusOK, gin.H{"visibility": "public"})
//     case http.StatusNotFound:
//         c.JSON(http.StatusNotFound, gin.H{"visibility": "private or non-existent"})
//     default:
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected response from GitHub API"})
//     }
// }

// func GetRepository(c *gin.Context) {

// }
