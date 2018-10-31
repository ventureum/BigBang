package lambda_feed_events_table_creator_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/feed/session_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/feed/purchase_mps_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/feed/refuel_record_config"
  "BigBang/internal/platform/postgres_config/feed/wallet_address_record_config"
)

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
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
  actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresBigBangClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresBigBangClient}
  purchaseMPsRecordExecutor := purchase_mps_record_config.PurchaseMPsRecordExecutor{*postgresBigBangClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresBigBangClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresBigBangClient}
  refuelRecordExecutor := refuel_record_config.RefuelRecordExecutor{*postgresBigBangClient}
  walletAddressRecordExecutor := wallet_address_record_config.WalletAddressRecordExecutor{*postgresBigBangClient}

  walletAddressRecordExecutor.DeleteWalletAddressRecordTable()
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
  walletAddressRecordExecutor.CreateWalletAddressRecordTable()

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}
