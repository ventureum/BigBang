package main

import (
	"BigBang/cmd/lambda/TCR/batch_objectives/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_batch_objectives_config.Handler)
}
