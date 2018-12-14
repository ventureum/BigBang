package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/delete_batch_actors/config"
)

func main() {
  lambda.Start(lambda_delete_batch_actors_config.Handler)
}
