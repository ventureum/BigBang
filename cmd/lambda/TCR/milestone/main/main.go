package main

import (
	"BigBang/cmd/lambda/TCR/milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_milestone_config.Handler)
}
