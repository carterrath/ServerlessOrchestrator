package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: 1, Name: "Item 1"},
	{ID: 2, Name: "Item 2"},
	// Add more items as needed
}

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

// RegisterRoutes registers API endpoints with the provided Gin router
func RegisterRoutes(router *gin.Engine) {
	router.GET("/items", GetItems)
	router.POST("/items", AddItem)
}

// APIStart initializes and starts the API server
func APIStart() {
	router := gin.Default()

	// Register API routes
	RegisterRoutes(router)

	// Run the server on port 8080
	router.Run(":8080")
}
