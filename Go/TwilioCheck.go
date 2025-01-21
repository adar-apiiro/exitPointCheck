package main

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/client"
	"github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/twilio-labs/sample-twilio-go/pkg/message"
)

func CreateMessage(client *twilio.RestClient, from string, to string, body string) {
	params := &v2010.CreateMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetBody(body)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent! SID: %s\n", *resp.Sid)
}

func main() {
	// Load Twilio credentials from environment variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSid == "" || authToken == "" {
		log.Fatal("TWILIO_ACCOUNT_SID or TWILIO_AUTH_TOKEN is not set")
	}

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(client.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Twilio phone number
	from := "+1234567890" // Replace with your Twilio phone number
	// Recipient phone number
	to := "+0987654321" // Replace with the recipient's phone number
	// Message body
	body := "Hello from Twilio using Go!"

	// Send the message
	CreateMessage(client, from, to, body)
}
