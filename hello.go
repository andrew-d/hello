package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, i Input) (string, error) {
	greeting := i.Name
	if greeting == "" {
		greeting = "world"
	}

	return fmt.Sprintf("Hello, %s!\n", greeting), nil
}

func main() {
	log.Println("started")
	lambda.Start(handler)
}
