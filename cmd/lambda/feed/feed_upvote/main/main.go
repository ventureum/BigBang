package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_upvote/config"
)

func main() {
  lambda.Start(lambda_feed_upvote_config.Handler)
}
