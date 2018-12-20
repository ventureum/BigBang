package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_validator_recent_rating_vote_activities/config"
)

func main() {
  lambda.Start(lambda_get_validator_recent_rating_vote_activities_config.Handler)
}
