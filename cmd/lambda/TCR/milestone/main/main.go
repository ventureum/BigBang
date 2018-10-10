package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/milestone/config"
)

func main() {
  lambda.Start(lambda_milestone_config.Handler)
}
