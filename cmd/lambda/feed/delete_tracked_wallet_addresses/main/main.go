package main

import (
	"BigBang/cmd/lambda/feed/delete_tracked_wallet_addresses/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_tracked_wallet_addresses_config.Handler)
}
