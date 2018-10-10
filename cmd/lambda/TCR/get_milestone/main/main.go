package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_milestone/config"
)

func main() {
  lambda.Start(lambda_get_milestone_config.Handler)
}
