package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
)

type Request struct {
  PostHash string `json:"postHash,required"`
}

type Response struct {
  Session *session_record_config.SessionRecordResult `json:"session,omitempty"`
  Ok bool `json:"ok"`
  Message string `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      response.Message = errStr.(string)
    }
  }()

  postHash := request.PostHash
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer postgresFeedClient.Close()
  postgresFeedClient.Begin()

  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}

  response.Session = sessionRecordExecutor.GetSessionRecordTx(postHash).ToSessionRecordResult()

  postgresFeedClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}


func main() {
  // TODO(david.shao): remove example when deployed to production
  //request := Request{
  //  PostHash: "0xpostHash001",
  //}
  //response, _ := Handler(request)
  //fmt.Printf("%+v", response)

  lambda.Start(Handler)
}
