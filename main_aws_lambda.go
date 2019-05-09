// +build aws_lambda

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	handler := Main()

	lambdaHandler := httpadapter.New(handler)

	lambda.Start(lambdaHandler.Proxy)
}
