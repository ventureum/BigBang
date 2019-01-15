package main

import (
	"BigBang/cmd/lambda/TCR/delete_objective/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_objective_config.Handler)
}
