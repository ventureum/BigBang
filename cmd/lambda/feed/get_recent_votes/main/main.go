package main

import (
	"BigBang/cmd/lambda/feed/get_recent_votes/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_recent_votes_config.Handler)
}
