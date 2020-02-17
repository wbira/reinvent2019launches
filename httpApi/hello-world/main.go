package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	content, err := json.Marshal(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %v", err),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Event: \n %v", string(content)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
