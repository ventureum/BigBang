package main

import (
	"BigBang/cmd/lambda/TCR/delete_milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_milestone_config.Handler)
}
