package main

import (
	"BigBang/cmd/lambda/TCR/delete_batch_objectives/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_batch_objectives_config.Handler)
}
