package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/attach_session/config"
)

func main() {
  lambda.Start(lambda_attach_session_config.Handler)
}
