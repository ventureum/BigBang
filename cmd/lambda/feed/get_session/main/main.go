package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_session/config"
)

func main() {
  lambda.Start(lambda_get_session_config.Handler)
}
