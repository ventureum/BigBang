package main

import (
	"BigBang/cmd/lambda/feed/get_next_redeem/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_next_redeem_config.Handler)
}
