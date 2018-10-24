package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/add_proxy/config"
)

func main() {
  lambda.Start(lambda_add_proxy_config.Handler)
}
