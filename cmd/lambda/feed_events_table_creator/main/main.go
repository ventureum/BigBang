package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed_events_table_creator/config"
)

func main() {
  lambda.Start(config.Handler)
}
