package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Varma1506/aws-crud-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ValidateCreateRequestBody(body models.CreateRequestBody) error {
	log.Print(body.Brand, body.Category, body.Description)
	if body.Brand == "" || body.Category == "" || body.Description == "" {
		return fmt.Errorf("Please Provide all values")
	}
	return nil
}

func SendMessageToSQS(message models.CreateRequestBody) error {
	log.Print("Sending Message to Queue")
	queueURL := os.Getenv("NOTIFY_QUEUE_URL")
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	sqsSvc := sqs.New(sess)

	// Convert the message to a JSON string
	jsonString, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	_, err = sqsSvc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(string(jsonString)),
	})

	if err != nil {
		// handle error
		return err
	}
	log.Println("Message Sent to Queue")
	return nil
}
