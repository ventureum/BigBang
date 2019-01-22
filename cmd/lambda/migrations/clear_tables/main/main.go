package main

import (
	"BigBang/cmd/lambda/migrations/clear_tables/config"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(clear_tables_config.Handler)
}
