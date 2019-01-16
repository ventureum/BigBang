package main

import (
	"BigBang/cmd/lambda/feed/feed_redeem_milestone_points/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_redeem_milestone_points_config.Handler)
}
