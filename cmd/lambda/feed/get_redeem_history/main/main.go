package main

import (
	"BigBang/cmd/lambda/feed/get_redeem_history/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_redeem_history_config.Handler)
}
