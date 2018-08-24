package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
  "BigBang/internal/pkg/error_config"
)


type Request struct {
  Actor string `json:"actor,required"`
}

type ResponseContent struct {
  Actor string `json:"actor"`
  ActorType string `json:"userType"`
  Reputations int64 `json:"reputations"`
}

type Response struct {
  Profile *ResponseContent `json:"profile,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProfileRecordResultToResponseContent(actorProfileRecord *actor_profile_record_config.ActorProfileRecord) *ResponseContent {
  return &ResponseContent{
    Actor: actorProfileRecord.Actor,
    ActorType: string(actorProfileRecord.ActorType),
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Profile = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()


  actor := request.Actor
  postgresFeedClient.Begin()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{*postgresFeedClient}

  actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecordTx(actor)

  response.Profile = ProfileRecordResultToResponseContent(actorProfileRecord)
  response.Profile.Reputations = actorReputationsRecordExecutor.GetActorReputationsTx(actor).Value()

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
  //  Actor: "0x001",
  //}
  //response, _ := Handler(request)
  //fmt.Printf("%+v\n", response)

  lambda.Start(Handler)
}
