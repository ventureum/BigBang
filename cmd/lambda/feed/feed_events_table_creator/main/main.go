package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/feed_events_table_creator/config"
)

func main() {
  lambda.Start(lambda_feed_events_table_creator_config.Handler)
}
