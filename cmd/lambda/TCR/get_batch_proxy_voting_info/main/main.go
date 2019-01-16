package main

import (
	"BigBang/cmd/lambda/TCR/get_batch_proxy_voting_info/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_batch_proxy_voting_info_config.Handler)
}
