package main

import (
	"BigBang/cmd/lambda/feed/attach_session/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_attach_session_config.Handler)
}
