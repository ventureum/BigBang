package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/adjust_proxy_votes/config"
)

func main() {
  lambda.Start(lambda_adjust_proxy_votes_config.Handler)
}
