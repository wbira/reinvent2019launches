package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
)

type SqsClient struct {
	client   *sqs.Client
	queueURL string
}

func (s *SqsClient) Init() error {
	queueURL := os.Getenv("QUEUE_URL")
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}
	client := sqs.New(cfg)
	s.client = client
	s.queueURL = queueURL
	return nil
}

func (s *SqsClient) SendMessege(groupId, messageId string) {
	input := &sqs.SendMessageInput{
		MessageGroupId:         aws.String(groupId),
		MessageDeduplicationId: aws.String(fmt.Sprintf("m-%v-%v", groupId, messageId)),
		MessageBody:            aws.String(messageId),
		QueueUrl:               aws.String(s.queueURL),
	}
	s.client.SendMessageRequest(input)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	client := SqsClient{}
	err := client.Init()
	if err != nil {
		return createResponse("Internal serverl error!", 500)
	}

	for i := 0; i < 10; i++ {
		client.SendMessege(string('A'), string(i))
	}
	return createResponse("Items sent successfuly!", 200)
}

func createResponse(body string, statusCode int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}
