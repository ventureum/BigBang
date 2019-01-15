package main

import (
	"BigBang/cmd/lambda/TCR/rating_vote/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_rating_vote_config.Handler)
}
