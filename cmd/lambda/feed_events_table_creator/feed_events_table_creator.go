package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/reputations_refuel_record_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/post_reputations_record_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
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

  db := client_config.ConnectPostgresClient()
  defer db.Close()

  db.Begin()
  db.LoadUuidExtension()
  db.LoadVoteTypeEnum()
  db.LoadActorTypeEnum()

  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*db}
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()

  reputationsRefuelRecordExecutor := reputations_refuel_record_config.ReputationsRefuelRecordExecutor{*db}
  reputationsRefuelRecordExecutor.DeleteReputationsRefuelRecordTable()
  reputationsRefuelRecordExecutor.CreateReputationsRefuelRecordTable()

  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*db}
  postRewardsRecordExecutor.DeletePostRewardsRecordTable()
  postRewardsRecordExecutor.CreatePostRewardsRecordTable()

  postReputationsRecordExecutor := post_reputations_record_config.PostReputationsRecordExecutor{*db}
  postReputationsRecordExecutor.DeletePostReputationsRecordTable()
  postReputationsRecordExecutor.CreatePostReputationsRecordTable()

  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*db}
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()

  postExecutor := post_config.PostExecutor{*db}
  postExecutor.DeletePostTable()
  postExecutor.CreatePostTable()

  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{*db}
  actorReputationsRecordExecutor.DeleteActorReputationsRecordTable()
  actorReputationsRecordExecutor.CreateActorReputationsRecordTable()

  db.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}


func main() {
  lambda.Start(Handler)
}
