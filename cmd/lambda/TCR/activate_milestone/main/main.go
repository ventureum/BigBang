package main

import (
	"BigBang/cmd/lambda/TCR/activate_milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_activate_milestone_config.Handler)
}
