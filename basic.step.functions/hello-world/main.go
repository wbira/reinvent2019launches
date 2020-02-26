package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Data struct {
	Comment string
}

type Event struct {
	Input Data
}

func handler(ctx context.Context, e Event) (string, error) {
	fmt.Println(e)
	fmt.Println(e.Input.Comment)
	return e.Input.Comment, nil
}

func main() {
	lambda.Start(handler)
}
