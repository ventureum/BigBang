package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/delete_project/config"
)

func main() {
  lambda.Start(lambda_delete_project_config.Handler)
}
