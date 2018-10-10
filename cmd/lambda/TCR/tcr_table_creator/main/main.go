package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/tcr_table_creator/config"
)

func main() {
  lambda.Start(lambda_tcr_table_creator_config.Handler)
}
