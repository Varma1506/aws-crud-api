package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Varma1506/aws-crud-api/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func notificationHandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	sess := session.Must(session.NewSession())
	snsClient := sns.New(sess)
	log.Println(sqsEvent.Records)
	log.Println("Check for first record Body")
	log.Println(sqsEvent.Records[0].Body)

	var userMessage models.CreateRequestBody
	message := sqsEvent.Records[0].Body
	log.Print("Notification Lambda is called")
	log.Println(sqsEvent.Records[0].Body)
	err := json.Unmarshal([]byte(message), &userMessage)
	if err != nil {
		return err
	}
	log.Print("Sending Mail")
	// Construct the SNS message
	snsMessage := fmt.Sprintf("Message to user %s: %s", userMessage.Brand, userMessage.Category)

	// Publish to SNS topic
	_, err = snsClient.Publish(&sns.PublishInput{
		Message:  aws.String(snsMessage),
		TopicArn: aws.String(os.Getenv("NOTIFY_TOPIC_ARN")),
	})
	if err != nil {
		fmt.Printf("Error publishing to SNS: %s\n", err)
		return err
	}
	return nil
}

func main() {
	lambda.Start(notificationHandler)
}
