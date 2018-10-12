package main

import (
  "github.com/aws/aws-lambda-go/lambda"
)

func main() {
  lambda.Start(lambda_rating_vote_config.Handler)
}
