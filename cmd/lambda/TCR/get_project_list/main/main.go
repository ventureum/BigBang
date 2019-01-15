package main

import (
	"BigBang/cmd/lambda/TCR/get_project_list/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_project_list_config.Handler)
}
