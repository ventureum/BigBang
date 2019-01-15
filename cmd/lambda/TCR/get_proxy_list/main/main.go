package main

import (
	"BigBang/cmd/lambda/TCR/get_proxy_list/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_proxy_list_config.Handler)
}
