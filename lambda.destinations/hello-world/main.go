package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Success bool
}

func handler(request MyEvent) error {
	fmt.Printf("Request success %v\n", request.Success)
	fmt.Printf("Request success type %T\n", request.Success)
	if request.Success {
		fmt.Println("Success path should applied")
		return nil
	} else {
		return errors.New("error, not success")
	}

}

func main() {
	lambda.Start(handler)
}
