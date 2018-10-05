package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/reset_actor_fuel/config"
)

func main() {
  lambda.Start(config.Handler)
}
