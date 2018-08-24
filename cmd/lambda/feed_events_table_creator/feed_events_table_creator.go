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
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/purchase_reputations_record_config"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/pkg/error_config"
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
  postgresFeedClient.SetIdleInTransactionSessionTimeout(60000)

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  actorProfileRecordExecutor.DeleteActorProfileRecordTable()
  actorProfileRecordExecutor.CreateActorProfileRecordTable()

  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{*postgresFeedClient}
  actorReputationsRecordExecutor.DeleteActorReputationsRecordTable()
  actorReputationsRecordExecutor.CreateActorReputationsRecordTable()

  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postExecutor.DeletePostTable()
  postExecutor.CreatePostTable()

  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()

  postReputationsRecordExecutor := post_reputations_record_config.PostReputationsRecordExecutor{*postgresFeedClient}
  postReputationsRecordExecutor.DeletePostReputationsRecordTable()
  postReputationsRecordExecutor.CreatePostReputationsRecordTable()

  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRewardsRecordExecutor.DeletePostRewardsRecordTable()
  postRewardsRecordExecutor.CreatePostRewardsRecordTable()

  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  postVotesCountersRecordExecutor.DeletePostVotesCountersRecordTable()
  postVotesCountersRecordExecutor.CreatePostVotesCountersRecordTable()

  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()

  purchaseReputationsRecordExecutor := purchase_reputations_record_config.PurchaseReputationsRecordExecutor{*postgresFeedClient}
  purchaseReputationsRecordExecutor.DeletePurchaseReputationsRecordTable()
  purchaseReputationsRecordExecutor.CreatePurchaseReputationsRecordTable()

  reputationsRefuelRecordExecutor := reputations_refuel_record_config.ReputationsRefuelRecordExecutor{*postgresFeedClient}
  reputationsRefuelRecordExecutor.DeleteReputationsRefuelRecordTable()
  reputationsRefuelRecordExecutor.CreateReputationsRefuelRecordTable()

  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  sessionRecordExecutor.DeleteSessionRecordTable()
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
