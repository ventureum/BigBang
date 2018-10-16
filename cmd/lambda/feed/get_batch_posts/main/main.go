package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/feed/get_batch_posts/config"
)

func main() {
  lambda.Start(lambda_get_batch_posts_config.Handler)
}
