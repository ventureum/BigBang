package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/refuel/config"
)

func main() {
  lambda.Start(config.Handler)
}
