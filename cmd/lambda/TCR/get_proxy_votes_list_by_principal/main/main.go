package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_proxy_votes_list_by_principal/config"
)

func main() {
  lambda.Start(lambda_get_proxy_voting_info_config.Handler)
}
