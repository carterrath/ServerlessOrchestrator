package api_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/api"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/suite"
)

// ApiTestSuite struct
type ApiTestSuite struct {
	suite.Suite
	router *gin.Engine
}

// SetupSuite runs before the test suite
func (suite *ApiTestSuite) SetupSuite() {
	// Initialize anything you need before running your test suite
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	api.RegisterRoutes(suite.router) // Register API routes
}

// func TestModelTestSuite(t *testing.T) {
// 	suite.Run(t, new(APITestSuite))
// }

// Test for retrieving items
func (suite *ApiTestSuite) TestGetItems() {
	// Create a new HTTP request to the GetItems endpoint
	req, _ := http.NewRequest(http.MethodGet, "/items", nil)
	resp := httptest.NewRecorder()

	// Serve the HTTP request using the router
	suite.router.ServeHTTP(resp, req)

	// Assert expectations using suite.Assertions
	suite.Equal(http.StatusOK, resp.Code, "Expected HTTP status code OK for GetItems")
	// Add more assertions as needed
}

// Test for adding a new item
func (suite *ApiTestSuite) TestAddItem() {
	// Test logic for adding a new item
	// Use httptest.NewRequest and suite.router.ServeHTTP to simulate a POST request
	// Assert expectations using suite.Assertions
}

// Run the test suite
func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}
