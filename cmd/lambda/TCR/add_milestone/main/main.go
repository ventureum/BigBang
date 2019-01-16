package main

import (
	"BigBang/cmd/lambda/TCR/add_milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_add_milestone_config.Handler)
}
