package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/finalize_milestone/config"
)

func main() {
  lambda.Start(lambda_finalize_milestone_config.Handler)
}
