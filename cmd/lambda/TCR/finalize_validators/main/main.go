package main

import (
	"BigBang/cmd/lambda/TCR/finalize_validators/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_finalize_validators_config.Handler)
}
