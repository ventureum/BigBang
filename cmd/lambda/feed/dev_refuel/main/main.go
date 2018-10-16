package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/dev_refuel/config"
)

func main() {
  lambda.Start(lambda_dev_refuel_config.Handler)
}
