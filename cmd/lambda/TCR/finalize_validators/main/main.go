package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/finalize_validators/config"
)

func main() {
  lambda.Start(lambda_finalize_validators_config.Handler)
}
