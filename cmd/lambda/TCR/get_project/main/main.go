package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_project/config"
)

func main() {
  lambda.Start(config.Handler)
}
