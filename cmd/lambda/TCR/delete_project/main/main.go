package main

import (
	"BigBang/cmd/lambda/TCR/delete_project/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_project_config.Handler)
}
