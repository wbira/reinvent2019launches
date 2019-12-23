package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(event interface{}) {
	fmt.Println(event)
}

func main() {
	lambda.Start(handler)
}
