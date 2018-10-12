package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/update_actor_rating_votes/config"
)

func main() {
  lambda.Start(lambda_update_actor_rating_votes_config.Handler)
}
