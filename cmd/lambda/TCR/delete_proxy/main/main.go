package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/delete_proxy/config"
)

func main() {
  lambda.Start(lambda_delete_proxy_config.Handler)
}
