package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/profile/config"
)

func main() {
  lambda.Start(lambda_profile_config.Handler)
}
