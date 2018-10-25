package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_actor_private_key/config"
)

func main() {
  lambda.Start(lambda_get_actor_private_key_config.Handler)
}
