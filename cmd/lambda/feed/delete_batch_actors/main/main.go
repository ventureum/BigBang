package main

import (
	"BigBang/cmd/lambda/feed/delete_batch_actors/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambda_delete_batch_actors_config.Handler)
}
