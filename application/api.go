package application

import (
	"net/http"

	"github.com/GoKubes/ServerlessOrchestrator/application/routes/microservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/routes/runmicroservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/routes/stopmicroservice"
	"github.com/GoKubes/ServerlessOrchestrator/application/routes/users"
	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/gin-gonic/gin"
)

func Init(dao *dataaccess.MicroservicesDAO, userdao *dataaccess.UserDAO, ecsClient *ecs.Client, r53Client *route53.Client) error {
	router := gin.Default()

	// CORS middleware first
	router.Use(corsMiddleware())

	// Other routes and middleware
	handleRoutes(router, dao, userdao, ecsClient, r53Client)

	// Run the server
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		return err
	}

	return nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "https://www.serverlessorchestrator.com" || origin == "https://serverlessorchestrator.com" || origin == "http://localhost:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func handleRoutes(router *gin.Engine, dao *dataaccess.MicroservicesDAO, userdao *dataaccess.UserDAO, ecsClient *ecs.Client, r53Client *route53.Client) {
	router.GET("/api/microservice", func(c *gin.Context) {
		microservice.GetAllMicroservices(c, dao, userdao)
	})
	router.POST("/api/microservice", func(c *gin.Context) {
		microservice.UploadMicroservice(c, dao)
	})
	router.POST("/api/signup/developer", func(c *gin.Context) {
		users.CreateDeveloper(c, userdao)
	})
	router.POST("/api/signup/consumer", func(c *gin.Context) {
		users.CreateConsumer(c, userdao)
	})
	router.POST("/api/login/developer", func(c *gin.Context) {
		users.Login(c, userdao)
	})
	router.POST("/api/login/consumer", func(c *gin.Context) {
		users.Login(c, userdao)
	})
	router.POST("/api/runmicroservice", func(c *gin.Context) {
		runmicroservice.RunMicroservice(c, dao, ecsClient, r53Client)
	})
	router.POST("/api/stopmicroservice", func(c *gin.Context) {
		stopmicroservice.StopMicroservice(c, dao)
	})
	router.GET("/api/getuserdetails", func(c *gin.Context) {
		users.GetUserDetails(c, userdao)
	})
	router.POST("/api/recovery", func(c *gin.Context) {
		users.Recovery(c, userdao)
	})
	router.POST("/api/verify-code", func(c *gin.Context) {
		users.Recovery(c, userdao)
	})
	router.POST("/api/reset", func(c *gin.Context) {
		users.ResetPassword(c, userdao)
	})
}
