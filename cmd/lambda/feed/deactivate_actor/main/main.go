package main

import (
	"BigBang/cmd/lambda/feed/deactivate_actor/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_deactivate_actor_config.Handler)
}
