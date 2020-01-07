package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.SQSEvent) error {
	for i, e := range request.Records {
		fmt.Printf("Messege id: %v, Body: %v, Source %v, index: %v\n", e.MessageId, e.Body, e.EventSource, i)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
