package main

import (
	"BigBang/cmd/lambda/TCR/delete_proxy/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_proxy_config.Handler)
}
