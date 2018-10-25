package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/set_actor_private_key/config"
)

func main() {
  lambda.Start(lambda_set_actor_private_key_config.Handler)
}
