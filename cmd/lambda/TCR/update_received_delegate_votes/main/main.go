package main

import (
	"BigBang/cmd/lambda/TCR/update_received_delegate_votes/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_update_received_delegate_votes_config.Handler)
}
