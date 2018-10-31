package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/add_tracked_wallet_addresses/config"
)

func main() {
  lambda.Start(lambda_add_tracked_wallet_addresses_config.Handler)
}
