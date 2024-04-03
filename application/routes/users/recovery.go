package users

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/GoKubes/ServerlessOrchestrator/dataaccess"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
)

// RecoveryRequest struct
type RecoveryRequest struct {
	Email string `json:"email" binding:"required"`
}

// Recovery function
func Recovery(c *gin.Context, userDAO *dataaccess.UserDAO) {
	var req RecoveryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email exists in the database
	_, err := userDAO.GetUserByEmail(req.Email)
	if err != nil {
		// If there's an error (user not found), return an error response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Generate a random 6-digit code
	code := generateRandomCode(100000, 999999)

	// Send recovery email with the code
	err = sendRecoveryEmail(req.Email, code)
	if err != nil {
		// If there's an error sending the email, return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send recovery email"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Recovery email sent"})
}

func generateRandomCode(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func sendRecoveryEmail(email string, code int) error {
	// Get AWS region from environment variable
	region := os.Getenv("AWS_REGION")

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}

	// Create an SES client
	svc := ses.New(sess)

	// Get sender email address from environment variable
	sender := os.Getenv("SENDER_EMAIL")

	// Specify the recipient email address
	recipient := email

	// Specify the email subject and body
	subject := "Password Recovery Code"
	body := fmt.Sprintf("Your password recovery code is: %d", code)

	// Create the email input
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	// Send the email
	_, err = svc.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}
