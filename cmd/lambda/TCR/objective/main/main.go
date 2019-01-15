package main

import (
	"BigBang/cmd/lambda/TCR/objective/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_objective_config.Handler)
}
