package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/purchase_mps_record_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/actor_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/refuel_record_config"
)

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()

  postgresFeedClient.Begin()
  postgresFeedClient.LoadUuidExtension()
  postgresFeedClient.LoadVoteTypeEnum()
  postgresFeedClient.LoadActorTypeEnum()
  postgresFeedClient.LoadActorProfileStatusEnum()
  postgresFeedClient.SetIdleInTransactionSessionTimeout(60000)

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresFeedClient}
  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}
  actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  purchaseMPsRecordExecutor := purchase_mps_record_config.PurchaseMPsRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{*postgresFeedClient}

  sessionRecordExecutor.DeleteSessionRecordTable()
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  purchaseMPsRecordExecutor.DeletePurchaseReputationsRecordTable()
  postVotesCountersRecordExecutor.DeletePostVotesCountersRecordTable()
  postRewardsRecordExecutor.DeletePostRewardsRecordTable()
  actorVotesCountersRecordExecutor.DeleteActorVotesCountersRecordTable()
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  refuelRecordExecutor.DeleteRefuelRecordTable()
  postExecutor.DeletePostTable()
  actorRewardsInfoRecordExecutor.DeleteActorRewardsInfoRecordTable()
  actorProfileRecordExecutor.DeleteActorProfileRecordTable()

  actorProfileRecordExecutor.CreateActorProfileRecordTable()
  actorRewardsInfoRecordExecutor.CreateActorRewardsInfoRecordTable()
  postExecutor.CreatePostTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()
  actorVotesCountersRecordExecutor.CreateActorVotesCountersRecordTable()
  postRewardsRecordExecutor.CreatePostRewardsRecordTable()
  postVotesCountersRecordExecutor.CreatePostVotesCountersRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()
  purchaseMPsRecordExecutor.CreatePurchaseMPsRecordTable()
  refuelRecordExecutor.CreateRefuelRecordTable()
  sessionRecordExecutor.CreateSessionRecordTable()

  postgresFeedClient.Commit()
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
