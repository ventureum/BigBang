package main

import (
  "log"
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/reputations_refuel_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
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

  reputationsRefuelRecord := reputations_refuel_record_config.ReputationsRefuelRecord{
    Actor: actor,
    Reputations: reputations,
  }

  postgresFeedClient.Begin()

  reputationsRefuelRecordExecutor := reputations_refuel_record_config.ReputationsRefuelRecordExecutor{
    *postgresFeedClient}
  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{
    *postgresFeedClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}


  actorProfileRecordExecutor.VerifyActorExistingTx(actor)
  actorReputationsRecordExecutor.VerifyActorExistingTx(actor)

  // Update Reputations Refuel Record
  reputationsRefuelRecordExecutor.UpsertReputationsRefuelRecordTx(&reputationsRefuelRecord)

  // Update Actor Reputations Record
  actorReputationsRecordExecutor.AddActorReputationsTx(actor, reputations)

  postgresFeedClient.Commit()

  log.Printf("Refueled %d to actor %s", reputations, actor)

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
