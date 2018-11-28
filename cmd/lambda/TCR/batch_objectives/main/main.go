package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/batch_objectives/config"
)

func main() {
  lambda.Start(lambda_batch_objectives_config.Handler)
}
