package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, e interface{}) {
	fmt.Println("Hello world!")
	fmt.Println(ctx)
	fmt.Println(e)
}

func main() {
	lambda.Start(handler)
}
