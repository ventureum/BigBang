package main

import (
	"BigBang/cmd/lambda/TCR/get_batch_rating_vote_list/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_batch_rating_vote_list_config.Handler)
}
