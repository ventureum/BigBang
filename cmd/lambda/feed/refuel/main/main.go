package main

import (
	"BigBang/cmd/lambda/feed/refuel/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_refuel_config.Handler)
}
