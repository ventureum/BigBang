package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/objective/config"
)

func main() {
  lambda.Start(lambda_objective_config.Handler)
}
