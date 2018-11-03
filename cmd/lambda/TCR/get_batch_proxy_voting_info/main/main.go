package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_batch_proxy_voting_info/config"
)

func main() {
  lambda.Start(lambda_get_batch_proxy_voting_info_config.Handler)
}
