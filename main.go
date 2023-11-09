package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleLambdaEvent(event Myevent) (MyResponse, error) {
	return MyResponse
}
func main() {
	lambda.Start(HandleLambdaEvent)
}
