package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
  "BigBang/internal/app/feed_attributes"
  "math"
)


type Request struct {
  Actor string `json:"actor,required"`
}

type ResponseContent struct {
  Actor string `json:"actor"`
  ActorType string `json:"actorType"`
  Level int64 `json:"level"`
  RewardsInfo *feed_attributes.RewardsInfo `json:"rewardsInfo"`
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
    }
    postgresFeedClient.Close()
  }()


  actor := request.Actor

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresFeedClient}

  actorProfileRecordExecutor.VerifyActorExisting(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExisting(actor)

  actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecord(actor)
  response.Profile = ProfileRecordResultToResponseContent(actorProfileRecord)
  log.Printf("Loaded Profile content for actor %s\n", actor)
  rewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actor)
  log.Printf("Loaded Rewards info for actor %s\n", actor)
  response.Profile.RewardsInfo = rewardsInfo
  response.Profile.Level = int64(math.Floor(math.Log10(1 + math.Max(float64(rewardsInfo.Reputation), 0))))
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
