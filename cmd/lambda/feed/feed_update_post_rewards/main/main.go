package main

import (
	"BigBang/cmd/lambda/feed/feed_update_post_rewards/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_update_post_rewards_config.Handler)
}
