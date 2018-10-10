package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_objective/config"
)

func main() {
  lambda.Start(lambda_get_objective_config.Handler)
}
