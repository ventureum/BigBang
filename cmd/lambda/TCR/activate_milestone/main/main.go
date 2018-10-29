package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/activate_milestone/config"
)

func main() {
  lambda.Start(lambda_activate_milestone_config.Handler)
}
