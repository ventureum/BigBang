package main

import (
	"BigBang/cmd/lambda/feed/feed_upvote/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_upvote_config.Handler)
}
