package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Varma1506/aws-crud-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Function to add new product to the dyanmoDB
func CreateNewProduct(body models.CreateRequestBody) (string, error) {
	log.Printf("Creating new product to the DB")
	id, err := generateUniqueID()
	//add id to the body object
	body.Id = id
	body.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	//Handle error
	if err != nil {
		return "", fmt.Errorf("Unable to generate Unique ID")
	}
	// Initialize a DynamoDB session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	svc := dynamodb.New(sess)

	// Prepare item for DynamoDB
	av, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		return "", err
	}

	// Put item in DynamoDB
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE_NAME")),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return "", err
	}

	err = SendMessageToSQS(body)

	if err != nil {
		return "", err
	}

	return id, nil
}

// function to generate unique ID for the record
func generateUniqueID() (string, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return string(newUUID), nil
}

//Function to get data from ID
