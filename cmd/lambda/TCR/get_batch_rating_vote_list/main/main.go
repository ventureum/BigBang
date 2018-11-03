package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_batch_rating_vote_list/config"
)

func main() {
  lambda.Start(lambda_get_batch_rating_vote_list_config.Handler)
}
