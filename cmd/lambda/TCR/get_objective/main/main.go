package main

import (
	"BigBang/cmd/lambda/TCR/get_objective/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_objective_config.Handler)
}
