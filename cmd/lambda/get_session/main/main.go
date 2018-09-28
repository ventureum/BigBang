package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/get_session/config"
)

func main() {
  lambda.Start(config.Handler)
}
