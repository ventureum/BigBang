package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/delete_milestone/config"
)

func main() {
  lambda.Start(lambda_delete_milestone_config.Handler)
}
