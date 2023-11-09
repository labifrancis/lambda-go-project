package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Struct in Go is basically creating your own data type. It contains different data types.
type Myevent struct {
	// Golang does not understand json that's why we are defining it like this.
	Name string `json: "what is your name?"`
	Age  int    `json: "How old are you"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func HandleLambdaEvent(event Myevent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}
func main() {
	lambda.Start(HandleLambdaEvent)
}
