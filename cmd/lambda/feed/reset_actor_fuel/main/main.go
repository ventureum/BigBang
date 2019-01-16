package main

import (
	"BigBang/cmd/lambda/feed/reset_actor_fuel/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_reset_actor_fuel_config.Handler)
}
