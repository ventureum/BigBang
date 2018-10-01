package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_recent_votes/config"
)

func main() {
  lambda.Start(config.Handler)
}
