package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Order struct {
	Id        string
	ProductId string
	Price     float64
}

type Event struct {
	Detail Order
}

func handler(event Event) error {
	order := event.Detail
	fmt.Printf("Order id %v\n", order.Id)
	fmt.Printf("Product id %v\n", order.ProductId)
	fmt.Printf("Price %v\n", order.Price)
	return nil
}

func main() {
	lambda.Start(handler)
}
