package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/add_proxy_voting_for_principal/config"
)

func main() {
  lambda.Start(lambda_add_proxy_voting_for_principal_config.Handler)
}
