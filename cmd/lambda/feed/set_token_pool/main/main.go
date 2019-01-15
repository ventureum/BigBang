package main

import (
	"BigBang/cmd/lambda/feed/set_token_pool/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_set_token_pool_config.Handler)
}
