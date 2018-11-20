package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/cmd/lambda/TCR/get_project_id_by_admin/config"
)

func main() {
  lambda.Start(lambda_get_project_id_by_admin_config.Handler)
}
