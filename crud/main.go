package main

import (
	"context"

	"github.com/Varma1506/aws-crud-api/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handlers.HandleGetProduct(request)
	case "POST":
		return handlers.HandleCreateProduct(request)
	case "PUT":
		return handlers.HandleUpdateProduct(request)
	case "DELETE":
		return handlers.HandleDeleteProduct(request)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
}

func main() {
	//Invoke the lambda
	lambda.Start(handler)
}
