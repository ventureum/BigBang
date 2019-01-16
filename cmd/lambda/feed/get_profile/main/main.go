package main

import (
	"BigBang/cmd/lambda/feed/get_profile/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_profile_config.Handler)
}
