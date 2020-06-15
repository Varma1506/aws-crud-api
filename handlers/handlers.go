package handlers

import (
	"encoding/json"
	"log"

	"github.com/Varma1506/aws-crud-api/helpers"
	"github.com/Varma1506/aws-crud-api/models"
	"github.com/aws/aws-lambda-go/events"
)

func HandleCreateProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body models.CreateRequestBody

	//Parse request body
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		// Handle error
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	//validate request body
	err = helpers.ValidateCreateRequestBody(body)
	if err != nil {
		//Handle Error
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	//Update the details to DB
	id, err := helpers.CreateNewProduct(body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       id,
	}, nil
}

func HandleGetProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Get request!!!")
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func HandleUpdateProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil

}

func HandleDeleteProduct(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil

}
