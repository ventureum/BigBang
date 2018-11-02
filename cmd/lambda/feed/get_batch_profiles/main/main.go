package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_batch_profiles/config"
)

func main() {
  lambda.Start(lambda_get_batch_profiles_config.Handler)
}
