package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  UserType string `json:"userType,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToActorProfileRecord() (*actor_profile_record_config.ActorProfileRecord) {
  return &actor_profile_record_config.ActorProfileRecord{
    Actor:      request.Actor,
    ActorType: feed_attributes.ActorType(request.UserType),
  }
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

  postgresFeedClient.Begin()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}

  inserted := actorProfileRecordExecutor.UpsertActorProfileRecordTx(request.ToActorProfileRecord())

  if inserted {
    actorReputationsRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
      *postgresFeedClient}
    actorReputationsRecord := actor_rewards_info_record_config.ActorRewardsInfoRecord{
      Actor:           request.Actor,
      Reputation:      feed_attributes.Reputation(feed_attributes.MuMinFuel.Value()),
      Fuel:            feed_attributes.MuMinFuel,
      MilestonePoints: 0,
    }
    actorReputationsRecordExecutor.UpsertActorRewardsInfoRecordTx(&actorReputationsRecord)

    log.Printf("Created Actor Fuel Account for actor %s", request.Actor)
  }

  postgresFeedClient.Commit()

  if inserted {
    log.Printf("Created Profile for actor %s", request.Actor)
  } else {
    log.Printf("Updated Profile for actor %s", request.Actor)
  }

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
