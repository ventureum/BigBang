package main

import (
	"BigBang/cmd/lambda/feed/feed_events_table_creator/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_events_table_creator_config.Handler)
}
