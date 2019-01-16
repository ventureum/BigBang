package main

import (
	"BigBang/cmd/lambda/feed/feed_post/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_post_config.Handler)
}
