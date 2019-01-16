package main

import (
	"BigBang/cmd/lambda/TCR/add_proxy/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_add_proxy_config.Handler)
}
