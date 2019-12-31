package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

type Customer struct {
	Address string
	Name    string
}

type Order struct {
	Id        string
	ProductId string
	Price     float64
	Customer  Customer
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var detailType string
	var source string
	var eventBusName string

	order := Order{
		Id:        "id",
		ProductId: "123213",
		Price:     float64(100.50),
		Customer:  Customer{Name: "Waldek", Address: "Wro"},
	}

	e, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	input := eventbridge.PutEventsInput{
		Entries: []eventbridge.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(e)),
				DetailType:   &detailType,
				EventBusName: &eventBusName,
				Source:       &source,
			},
		},
	}
	fmt.Println(input)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("event sent"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
