package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
)

type Response struct {
  Ok bool `json:"ok"`
  Message string `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      response.Message = errStr.(string)
    }
  }()

  postgresFeedClient := client_config.ConnectPostgresClient()
  defer postgresFeedClient.Close()

  postgresFeedClient.Begin()
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRewardsRecordExecutor.UpdatePostRewardsRecordsByAggregationsTx()
  postgresFeedClient.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}


func main() {
  // TODO(david.shao): remove example when deployed to production
  //response, _ := Handler()
  //fmt.Println("%+v\n", response.Ok)

  lambda.Start(Handler)
}
