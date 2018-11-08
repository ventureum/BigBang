package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_next_redeem/config"
)

func main() {
  lambda.Start(lambda_get_next_redeem_config.Handler)
}
