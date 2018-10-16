package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_token_generator/config"
)

func main() {
  lambda.Start(lambda_feed_token_generator_config.Handler)
}
