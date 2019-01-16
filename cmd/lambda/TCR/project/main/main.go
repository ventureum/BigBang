package main

import (
	"BigBang/cmd/lambda/TCR/project/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_project_config.Handler)
}
