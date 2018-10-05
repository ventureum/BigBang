package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/rating_vote/config"
)

func main() {
  lambda.Start(config.Handler)
}
