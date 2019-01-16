package main

import (
	"BigBang/cmd/lambda/feed/profile/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_profile_config.Handler)
}
