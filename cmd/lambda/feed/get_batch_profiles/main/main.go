package main

import (
	"BigBang/cmd/lambda/feed/get_batch_profiles/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_batch_profiles_config.Handler)
}
