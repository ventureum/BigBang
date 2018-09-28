package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/get_feed_post/config"
)

func main() {
  lambda.Start(config.Handler)
}
