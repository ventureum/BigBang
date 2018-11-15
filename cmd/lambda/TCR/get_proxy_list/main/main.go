package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_proxy_list/config"
)

func main() {
  lambda.Start(lambda_get_proxy_list_config.Handler)
}
