package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed_update_post_rewards/config"
)

func main() {
  lambda.Start(config.Handler)
}
