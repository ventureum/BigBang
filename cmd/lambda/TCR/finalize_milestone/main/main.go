package main

import (
	"BigBang/cmd/lambda/TCR/finalize_milestone/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_finalize_milestone_config.Handler)
}
