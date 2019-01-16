package main

import (
	"BigBang/cmd/lambda/TCR/get_milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_milestone_config.Handler)
}
