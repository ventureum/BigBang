package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/deactivate_actor/config"
)

func main() {
  lambda.Start(config.Handler)
}
