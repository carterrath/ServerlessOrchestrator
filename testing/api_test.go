package api_tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoKubes/ServerlessOrchestrator/Application"
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
	Application.RegisterRoutes(suite.router) // Register API routes
}

// TC-001: Database fetching for all microservices.
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

// TC-002: Upload a new microservice, with a public GithHub repository link.
func (suite *ApiTestSuite) TestPublicGitHubLink() {
	// Example GitHub API URL: "https://api.github.com/repos/%7Bowner%7D/%7Brepo%7D"
	testURL := "https://api.github.com/repos/octocat/Hello-World"
	// Make sure the above link is PUBLIC
	req, _ := http.NewRequest("GET", "/check-repo-visibility?repo_url="+testURL, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)

}

// TC-003: Upload a new microservice, with a private GithHub repository link.
func (suite *ApiTestSuite) TestPrivateGitHubLink() {
	// Example GitHub API URL: "https://api.github.com/repos/%7Bowner%7D/%7Brepo%7D"
	testURL := "https://api.github.com/repos/carterrath/PrivateRepo"
	req, _ := http.NewRequest("GET", "/check-repo-visibility?repo_url="+testURL, nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	print("w.code here:", w.Code)
	suite.Equal(http.StatusNotFound, w.Code)

}

//TC-004: Validate system behavior when there is no Docker Image on file for a requested microservice. (i.e. the microservice has not been built yet)

//TC-005: Response when the requested microservice does not work.

//TC-006: Assess system response when the microservice is not ecxecuted and displayed as expected.

//TC-007: Verify successful microservice deployment and output delivery to the user.

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
