package main

import (
	"BigBang/cmd/lambda/TCR/get_batch_finalized_validators/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_batch_finalized_validators_config.Handler)
}
