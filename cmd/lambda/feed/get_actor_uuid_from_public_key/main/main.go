package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_actor_uuid_from_public_key/config"
)

func main() {
  lambda.Start(lambda_get_actor_uuid_from_public_key_config.Handler)
}
