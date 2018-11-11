package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_redeem_milestone_points/config"
)

func main() {
  lambda.Start(lambda_feed_redeem_milestone_points_config.Handler)
}
