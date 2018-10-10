package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/delete_objective/config"
)

func main() {
  lambda.Start(lambda_delete_objective_config.Handler)
}
