package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_redeem_block_info/config"
)

func main() {
  lambda.Start(lambda_get_redeem_block_info_config.Handler)
}
