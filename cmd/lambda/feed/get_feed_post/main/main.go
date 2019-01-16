package main

import (
	"BigBang/cmd/lambda/feed/get_feed_post/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_feed_post_config.Handler)
}
