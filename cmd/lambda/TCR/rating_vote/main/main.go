package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/rating_vote/config"
)

func main() {
  lambda.Start(lambda_rating_vote_config.Handler)
}
