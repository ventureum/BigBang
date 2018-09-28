package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed_upvote/config"
)

func main() {
  lambda.Start(config.Handler)
}
