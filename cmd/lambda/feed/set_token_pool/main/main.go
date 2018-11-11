package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/set_token_pool/config"
)

func main() {
  lambda.Start(lambda_set_token_pool_config.Handler)
}
