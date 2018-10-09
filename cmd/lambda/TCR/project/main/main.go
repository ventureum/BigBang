package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/project/config"
)

func main() {
  lambda.Start(lambda_project_config.Handler)
}
