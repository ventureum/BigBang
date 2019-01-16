package main

import (
	"BigBang/cmd/lambda/TCR/add_proxy_voting_for_principal/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_add_proxy_voting_for_principal_config.Handler)
}
