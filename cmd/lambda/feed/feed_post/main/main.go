package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_post/config"
)

func main() {
  lambda.Start(lambda_feed_post_config.Handler)
}
