package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
)

type Request struct {
  PostHash string `json:"postHash,required"`
}

type Response struct {
  Session *session_record_config.SessionRecordResult `json:"session,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo`json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Session = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresFeedClient.Close()
  }()

  postHash := request.PostHash

  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  response.Session = sessionRecordExecutor.GetSessionRecord(postHash).ToSessionRecordResult()

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
