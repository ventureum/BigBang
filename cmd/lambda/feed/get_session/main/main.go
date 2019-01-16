package main

import (
	"BigBang/cmd/lambda/feed/get_session/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_session_config.Handler)
}
