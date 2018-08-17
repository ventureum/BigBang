package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  UserType string `json:"userType,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message string `json:"message,omitempty"`
}

func (request *Request) ToActorProfileRecord() (*actor_profile_record_config.ActorProfileRecord) {
  return &actor_profile_record_config.ActorProfileRecord{
    Actor:      request.Actor,
    ActorType: feed_attributes.ActorType(request.UserType),
  }
}

func ProcessRequest(request Request, response *Response) {
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      response.Message = errStr.(string)
    }
  }()

  postgresFeedClient := client_config.ConnectPostgresClient()
  defer postgresFeedClient.Close()

  postgresFeedClient.Begin()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorProfileRecordExecutor.UpsertActorProfileRecordTx(request.ToActorProfileRecord())

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
  // Actor:  "0x005",
  // UserType: "KOL",
  //}
  //response, _ := Handler(request)
  //log.Printf("%+v\n",  response)

  lambda.Start(Handler)
}
