package main

import (
	"BigBang/cmd/lambda/TCR/tcr_table_creator/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_tcr_table_creator_config.Handler)
}
