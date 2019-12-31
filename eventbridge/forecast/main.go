package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) error {
	fmt.Println(event)
	return nil
}

func main() {
	lambda.Start(handler)
}
