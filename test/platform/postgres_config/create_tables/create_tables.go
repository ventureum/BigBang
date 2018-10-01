package main

import (
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/client_config"
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
  postgresBigBangClient := client_config.ConnectPostgresClient()

  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      message := error_config.CreatedErrorInfoFromString(errPanic)
      log.Printf("Error: %+v", message)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  postgresBigBangClient.Begin()
  postgresBigBangClient.LoadUuidExtension()
  postgresBigBangClient.LoadVoteTypeEnum()
  postgresBigBangClient.LoadActorTypeEnum()
  postgresBigBangClient.LoadActorProfileStatusEnum()
  postgresBigBangClient.SetIdleInTransactionSessionTimeout(60000)

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{*postgresBigBangClient}
  postExecutor := post_config.PostExecutor{*postgresBigBangClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresBigBangClient}
  postReputationsRecordExecutor := actor_vote_counter_record_config.PostReputationsRecordExecutor{*postgresBigBangClient}
  postMPsRecordExecutor := post_rewards_record_config.PostMPsRecordExecutor{*postgresBigBangClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresBigBangClient}
  purchaseMPsRecordExecutor := purchase_mps_record_config.PurchaseMPsRecordExecutor{*postgresBigBangClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresBigBangClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresBigBangClient}
  fuelsRefuelRecordExecutor := refuel_record_config.FuelsRefuelRecordExecutor{*postgresBigBangClient}

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

  postgresBigBangClient.Commit()
}