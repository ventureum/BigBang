package main

import (
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/client_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_reputations_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/feed/session_record_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/fuels_refuel_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_mps_record_config"
  "BigBang/internal/platform/postgres_config/feed/purchase_mps_record_config"
)


func main() {
  postgresFeedClient := client_config.ConnectPostgresClient()

  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      log.Printf("Error: %+v", message)
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
  postReputationsRecordExecutor := actor_vote_counter_record_config.PostReputationsRecordExecutor{*postgresFeedClient}
  postMPsRecordExecutor := post_rewards_record_config.PostMPsRecordExecutor{*postgresFeedClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  purchaseMPsRecordExecutor := purchase_mps_record_config.PurchaseMPsRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  fuelsRefuelRecordExecutor := refuel_record_config.FuelsRefuelRecordExecutor{*postgresFeedClient}

  sessionRecordExecutor.DeleteSessionRecordTable()
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  purchaseMPsRecordExecutor.DeletePurchaseReputationsRecordTable()
  postVotesCountersRecordExecutor.DeletePostVotesCountersRecordTable()
  postMPsRecordExecutor.DeletePostMPsRecordTable()
  postReputationsRecordExecutor.DeletePostReputationsRecordTable()
  postRepliesRecordExecutor.DeletePostRepliesRecordTable()
  fuelsRefuelRecordExecutor.DeleteFuelsRefuelRecordTable()
  postExecutor.DeletePostTable()
  actorRewardsInfoRecordExecutor.DeleteActorRewardsInfoRecordTable()
  actorProfileRecordExecutor.DeleteActorProfileRecordTable()

  actorProfileRecordExecutor.CreateActorProfileRecordTable()
  actorRewardsInfoRecordExecutor.CreateActorRewardsInfoRecordTable()
  postExecutor.CreatePostTable()
  postRepliesRecordExecutor.CreatePostRepliesRecordTable()
  postReputationsRecordExecutor.CreatePostReputationsRecordTable()
  postMPsRecordExecutor.CreatePostMPsRecordTable()
  postVotesCountersRecordExecutor.CreatePostVotesCountersRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()
  purchaseMPsRecordExecutor.CreatePurchaseMPsRecordTable()
  fuelsRefuelRecordExecutor.CreateFuelsRefuelRecordTable()
  sessionRecordExecutor.CreateSessionRecordTable()

  postgresFeedClient.Commit()
}