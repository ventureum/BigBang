package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_recent_votes/config"
)

func main() {
  lambda.Start(lambda_get_recent_votes_config.Handler)
}
