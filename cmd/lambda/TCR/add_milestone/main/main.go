package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/add_milestone/config"
)

func main() {
  lambda.Start(lambda_add_milestone_config.Handler)
}
