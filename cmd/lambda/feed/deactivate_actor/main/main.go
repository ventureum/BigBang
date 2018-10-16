package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/deactivate_actor/config"
)

func main() {
  lambda.Start(lambda_deactivate_actor_config.Handler)
}
