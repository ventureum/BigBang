package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/reset_actor_fuel/config"
)

func main() {
  lambda.Start(lambda_reset_actor_fuel_config.Handler)
}
