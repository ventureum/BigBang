package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_update_post_rewards/config"
)

func main() {
  lambda.Start(lambda_feed_update_post_rewards_config.Handler)
}
