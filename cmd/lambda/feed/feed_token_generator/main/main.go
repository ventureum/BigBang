package main

import (
	"BigBang/cmd/lambda/feed/feed_token_generator/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_feed_token_generator_config.Handler)
}
