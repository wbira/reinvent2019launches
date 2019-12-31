package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
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

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	eventBusName := os.Getenv("EVENT_BUS_NAME")

	order, err := getMarshalOrder()
	if err != nil {
		fmt.Printf("Cant marshall order %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	input := eventbridge.PutEventsInput{
		Entries: []eventbridge.PutEventsRequestEntry{
			{
				Detail:       aws.String(string(order)),
				DetailType:   aws.String("ORDER_CREATED"),
				EventBusName: &eventBusName,
				Source:       aws.String(string("Order service")),
			},
		},
	}

	client, err := initializeEventBridgeClient()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	if _, err := client.PutEventsRequest(&input).Send(ctx); err != nil {
		fmt.Printf("Cant put event %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("event sent"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func initializeEventBridgeClient() (*eventbridge.Client, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Printf("Cant load config %v", err)
		return nil, err
	}
	return eventbridge.New(cfg), nil
}

func getMarshalOrder() ([]byte, error) {
	order := Order{
		Id:        "id",
		ProductId: "123213",
		Price:     float64(100.50),
		Customer:  Customer{Name: "Waldek", Address: "Wro"},
	}
	return json.Marshal(order)
}
