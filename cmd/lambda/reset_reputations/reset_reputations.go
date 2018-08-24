package main

import (
  "log"
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
  "BigBang/internal/pkg/error_config"
)

type Request struct {
  UserAddress string `json:"userAddress,required"`
  Reputations int64 `json:"reputations,required"`
}

type Response struct {
  Ok      bool   `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()

  reputations := feed_attributes.Reputation(request.Reputations)
  actor := request.UserAddress

  postgresFeedClient.Begin()

  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{
    *postgresFeedClient}

  actorReputationsRecord := actor_reputations_record_config.ActorReputationsRecord{
    Actor: actor,
    Reputations: reputations,
  }

  actorReputationsRecordExecutor.UpsertActorReputationsRecordTx(&actorReputationsRecord)

  postgresFeedClient.Commit()

  log.Printf("Reset reputations for actor %s to be %d", actor, reputations)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}

func main() {
  //TODO(david.shao): remove example when deployed to production
  //request := Request{
  // UserAddress: "0x003",
  // Reputations: 400000,
  //}
  //Handler(request)

  lambda.Start(Handler)
}
