package main

import (
	"context"
	"log"

	"github.com/Varma1506/aws-crud-api/helpers"
	"github.com/Varma1506/aws-crud-api/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func cronJobHandler(ctx context.Context, cronEvent events.CloudWatchEvent) error {

	resoruces := cronEvent.Resources

	log.Printf("This is the Event logged from Cron Job : ")
	log.Print(resoruces)

	var body models.CreateRequestBody

	body.Brand = helpers.SetRandomBrand()
	body.Category = "Phone"
	body.Description = helpers.SetRandomDescription()
	body.Price = helpers.SetRandomPrice()

	//Add these random values to DB
	id, err := helpers.CreateNewProduct(body)
	if err != nil {
		log.Printf(err.Error(), id)
	}
	return nil
}

func main() {
	lambda.Start(cronJobHandler)
}
