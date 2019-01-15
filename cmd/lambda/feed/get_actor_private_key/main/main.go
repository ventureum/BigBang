package main

import (
	"BigBang/cmd/lambda/feed/get_actor_private_key/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_get_actor_private_key_config.Handler)
}
