package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/app/feed_attributes"
)


type Request struct {
  Actor string `json:"actor,required"`
  TypeHash string `json:"typeHash,required"`
  Limit int64 `json:"limit,omitempty"`
}

type Response struct {
  RecentPosts *[]post_rewards_record_config.PostRewardsRecord `json:"recentPosts,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.RecentPosts = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresFeedClient.Close()
  }()

  actor := request.Actor
  postType := feed_attributes.CreatePostTypeFromHashStr(request.TypeHash)
  limit := request.Limit

  if limit == 0 {
    limit = 20
  }

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresFeedClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}

  actorProfileRecordExecutor.VerifyActorExisting(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExisting(actor)

  response.RecentPosts =  postRewardsRecordExecutor.GetRecentPostRewardsRecordsByActor(actor, postType, limit)

  log.Printf("RecentPostRewardsRecords is loaded for actor %s\n", actor)

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
  // Actor: "0x009",
  //}
  //response, _ := Handler(request)
  //fmt.Printf("%+v", response)

  lambda.Start(Handler)
}
