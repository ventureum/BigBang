package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/set_next_redeem/config"
)

func main() {
  lambda.Start(lambda_set_next_redeem_config.Handler)
}
