package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/profile/config"
)

func main() {
  lambda.Start(config.Handler)
}
