package clear_tables_config

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_validator_record_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
	"BigBang/internal/platform/postgres_config/TCR/proxy_config"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_milestone_points_redeem_history_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_votes_counters_record_config"
	"BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
	"BigBang/internal/platform/postgres_config/feed/purchase_mps_record_config"
	"BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/refuel_record_config"
	"BigBang/internal/platform/postgres_config/feed/session_record_config"
	"BigBang/internal/platform/postgres_config/feed/wallet_address_record_config"
)

type Request struct {
	DBInfo *client_config.DBInfo `json:"dbInfo,omitempty"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(request.DBInfo)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()

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

	milestonePointsRedeemRequestRecordExecutor := milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{*postgresBigBangClient}
	redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}
	actorMilestonePointsRedeemHistoryRecordExecutor := actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecordExecutor{*postgresBigBangClient}
	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
	proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
	actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
	principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}
	milestoneValidatorRecordExecutor := milestone_validator_record_config.MilestoneValidatorRecordExecutor{*postgresBigBangClient}


	principalProxyVotesExecutor.ClearPrincipalProxyVotesTable()
	milestoneValidatorRecordExecutor.ClearMilestoneValidatorRecordTable()
	actorDelegateVotesAccountExecutor.ClearActorRatingVoteAccountTable()
	objectiveExecutor.ClearObjectiveTable()
	milestoneExecutor.ClearMilestoneTable()
	projectExecutor.ClearProjectTable()
	proxyExecutor.ClearProxyTable()
	ratingVoteExecutor.ClearRatingVoteTable()
	actorMilestonePointsRedeemHistoryRecordExecutor.ClearActorMilestonePointsRedeemHistoryRecordTable()
	redeemBlockInfoRecordExecutor.ClearRedeemBlockInfoRecordTable()
	milestonePointsRedeemRequestRecordExecutor.ClearMilestonePointsRedeemRequestRecordTable()

	walletAddressRecordExecutor.ClearWalletAddressRecordTable()
	sessionRecordExecutor.ClearSessionRecordTable()
	postVotesRecordExecutor.ClearPostVotesRecordTable()
	purchaseMPsRecordExecutor.ClearPurchaseReputationsRecordTable()
	postVotesCountersRecordExecutor.ClearPostVotesCountersRecordTable()
	postRewardsRecordExecutor.ClearPostRewardsRecordTable()
	actorVotesCountersRecordExecutor.ClearActorVotesCountersRecordTable()
	postRepliesRecordExecutor.ClearPostRepliesRecordTable()
	refuelRecordExecutor.ClearRefuelRecordTable()
	postExecutor.ClearPostTable()
	actorRewardsInfoRecordExecutor.ClearActorRewardsInfoRecordTable()
	actorProfileRecordExecutor.ClearActorProfileRecordTable()

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
