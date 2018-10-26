package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/update_available_delegate_votes/config"
)

func main() {
  lambda.Start(lambda_update_available_delegate_votes_config.Handler)
}
