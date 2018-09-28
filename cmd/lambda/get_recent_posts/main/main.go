package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/get_recent_posts/config"
)

func main() {
  lambda.Start(config.Handler)
}
