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
  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{*postgresFeedClient}
  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}
  postReputationsRecordExecutor := post_reputations_record_config.PostReputationsRecordExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  purchaseReputationsRecordExecutor := purchase_reputations_record_config.PurchaseReputationsRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  reputationsRefuelRecordExecutor := reputations_refuel_record_config.ReputationsRefuelRecordExecutor{*postgresFeedClient}

  sessionRecordExecutor.DeleteSessionRecordTable()
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  purchaseReputationsRecordExecutor.DeletePurchaseReputationsRecordTable()
  postVotesCountersRecordExecutor.DeletePostVotesCountersRecordTable()
  postRewardsRecordExecutor.DeletePostRewardsRecordTable()
  postReputationsRecordExecutor.DeletePostReputationsRecordTable()
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  reputationsRefuelRecordExecutor.DeleteReputationsRefuelRecordTable()
  postExecutor.DeletePostTable()
  actorReputationsRecordExecutor.DeleteActorReputationsRecordTable()
  actorProfileRecordExecutor.DeleteActorProfileRecordTable()


  actorProfileRecordExecutor.CreateActorProfileRecordTable()
  actorReputationsRecordExecutor.CreateActorReputationsRecordTable()
  postExecutor.CreatePostTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()
  postReputationsRecordExecutor.CreatePostReputationsRecordTable()
  postRewardsRecordExecutor.CreatePostRewardsRecordTable()
  postVotesCountersRecordExecutor.CreatePostVotesCountersRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()
  purchaseReputationsRecordExecutor.CreatePurchaseReputationsRecordTable()
  reputationsRefuelRecordExecutor.CreateReputationsRefuelRecordTable()
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
