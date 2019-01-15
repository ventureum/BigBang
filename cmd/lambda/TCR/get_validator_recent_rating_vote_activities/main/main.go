package main

import (
	"BigBang/cmd/lambda/TCR/get_validator_recent_rating_vote_activities/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_validator_recent_rating_vote_activities_config.Handler)
}
