package main

import (
	"BigBang/cmd/lambda/feed/dev_refuel/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_dev_refuel_config.Handler)
}
