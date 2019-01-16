package lambda_feed_upvote_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_votes_counters_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
	"log"
	"math"
	"time"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor    string `json:"actor,required"`
	PostHash string `json:"postHash,required"`
	Value    int64  `json:"value,required"`
}

type Response struct {
	VoteInfo *feed_attributes.VoteInfo `json:"voteInfo,omitempty"`
	Ok       bool                      `json:"ok"`
	Message  *error_config.ErrorInfo   `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.VoteInfo = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			if (feed_attributes.CreateVoteTypeFromValue(request.Body.Value) != feed_attributes.LOOKUP_VOTE_TYPE) &&
				postgresBigBangClient.Tx != nil {
				postgresBigBangClient.RollBack()
			}
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	postVotesRecord := post_votes_record_config.PostVotesRecord{
		Actor:    request.Body.Actor,
		PostHash: request.Body.PostHash,
		VoteType: feed_attributes.CreateVoteTypeFromValue(request.Body.Value),
	}

	if postVotesRecord.VoteType == feed_attributes.LOOKUP_VOTE_TYPE {
		response.VoteInfo = QueryPostVotesInfo(&postVotesRecord, postgresBigBangClient)
	} else {
		response.VoteInfo = ProcessPostVotesRecord(&postVotesRecord, postgresBigBangClient)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func QueryPostVotesInfo(
	postVotesRecord *post_votes_record_config.PostVotesRecord,
	postgresBigBangClient *client_config.PostgresBigBangClient) *feed_attributes.VoteInfo {
	var voteInfo feed_attributes.VoteInfo

	actor := postVotesRecord.Actor
	postHash := postVotesRecord.PostHash

	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	postExecutor := post_config.PostExecutor{*postgresBigBangClient}
	postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{
		*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
	actorRewardsInfoRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
	postExecutor.VerifyPostRecordExistingTx(postVotesRecord.PostHash)

	// Current Actor ActualMilestonePoints
	actorRewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)
	log.Printf("Current Actor RewardsInfo: %+v\n", actorRewardsInfo)
	voteInfo.RewardsInfo = actorRewardsInfo

	// Total Actor Reputation
	totalActorReputations := actorRewardsInfoRecordExecutor.GetTotalActorReputationTx()

	log.Printf("Total Actor Reputation: %+v\n", totalActorReputations)

	postVotesCountersRecord := postVotesCountersRecordExecutor.GetPostVotesCountersRecordByPostHashTx(postHash)

	// Total Actor Reputation for PostHash
	totalReputationsForPostHash := actorVotesCountersRecordExecutor.GetTotalActorReputationByPostHashTx(postHash)

	log.Printf("Total Actor Reputation for PostHash %s: %+v\n", postHash, totalReputationsForPostHash)

	// Calculate FuelCost
	var fuelCost feed_attributes.Fuel
	if totalActorReputations > 0 {
		fuelCost = feed_attributes.Fuel(math.Round(float64(feed_attributes.BetaMax) *
			(1.00 - float64(totalReputationsForPostHash)/(float64(totalActorReputations)))))
	} else {
		fuelCost = feed_attributes.BetaMax
	}
	voteInfo.FuelCost = feed_attributes.Fuel(fuelCost)
	log.Printf("FuelCost for PostHash %s: %+v\n", postHash, voteInfo.FuelCost)

	actorVotesCountersRecord := actorVotesCountersRecordExecutor.GetActorVotesCountersRecordByPostHashAndActorTx(
		postHash, actor)

	voteInfo.PostHash = postVotesRecord.PostHash
	voteInfo.Actor = postVotesRecord.Actor
	voteInfo.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
		UpVoteCount:    postVotesCountersRecord.UpVoteCount,
		DownVoteCount:  postVotesCountersRecord.DownVoteCount,
		TotalVoteCount: postVotesCountersRecord.TotalVoteCount,
	}
	voteInfo.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
		UpVoteCount:    actorVotesCountersRecord.UpVoteCount,
		DownVoteCount:  actorVotesCountersRecord.DownVoteCount,
		TotalVoteCount: actorVotesCountersRecord.TotalVoteCount,
	}

	return &voteInfo
}

func ProcessPostVotesRecord(
	postVotesRecord *post_votes_record_config.PostVotesRecord,
	postgresBigBangClient *client_config.PostgresBigBangClient) *feed_attributes.VoteInfo {

	actor := postVotesRecord.Actor
	postHash := postVotesRecord.PostHash
	voteType := postVotesRecord.VoteType

	var voteInfo feed_attributes.VoteInfo
	voteInfo.Actor = postVotesRecord.Actor
	voteInfo.PostHash = postVotesRecord.PostHash

	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{
		*postgresBigBangClient}
	postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{
		*postgresBigBangClient}
	postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	postExecutor := post_config.PostExecutor{*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
	actorRewardsInfoRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
	postExecutor.VerifyPostRecordExistingTx(postVotesRecord.PostHash)

	// CutOff Time
	cutOffTimeStamp := time.Now()

	// Actor List for PostHash and VoteType

	var actorList []string

	actorList = *postVotesRecordExecutor.GetActorListByPostHashAndVoteTypeTx(
		postVotesRecord.PostHash, postVotesRecord.VoteType)
	log.Printf("Actor List for PostHash and VoteType: %+v\n", actorList)

	// Current Actor ActualMilestonePoints
	actorRewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)
	log.Printf("Current Actor RewardsInfo: %+v\n", actorRewardsInfo)

	// Total Actor Reputation
	totalActorReputations := actorRewardsInfoRecordExecutor.GetTotalActorReputationTx()
	log.Printf("Total Actor Reputation: %+v\n", totalActorReputations)

	postVotesCountersRecord := postVotesCountersRecordExecutor.GetPostVotesCountersRecordByPostHashTx(postHash)

	// Total Actor Reputation for PostHash
	totalReputationsForPostHash := actorVotesCountersRecordExecutor.GetTotalActorReputationByPostHashTx(postHash)
	log.Printf("Total Actor Reputation for PostHash %s: %+v\n", postHash, totalReputationsForPostHash)

	// Total Reputation for PostHash with the same voteType as actor
	var totalReputationsForPostHashWithSameVoteType feed_attributes.Reputation
	if voteType == feed_attributes.UP_VOTE_TYPE {
		totalReputationsForPostHashWithSameVoteType = postVotesCountersRecord.TotalReputationForUpvote
	} else {
		totalReputationsForPostHashWithSameVoteType = postVotesCountersRecord.TotalReputationForDownvote
	}
	log.Printf("Total Actor Reputaions for PostHash with the same voteType as actor: %+v\n",
		totalReputationsForPostHashWithSameVoteType)

	// Calculate FuelCost
	var fuelCost feed_attributes.Fuel
	if totalActorReputations > 0 {
		fuelCost = feed_attributes.Fuel(math.Round(float64(feed_attributes.BetaMax) *
			(1.00 - float64(totalReputationsForPostHash)/(float64(totalActorReputations)))))
	} else {
		fuelCost = feed_attributes.BetaMax
	}
	voteInfo.FuelCost = feed_attributes.Fuel(fuelCost)
	log.Printf("FuelCost for PostHash %s: %+v\n", postHash, voteInfo.FuelCost)

	// Update Fuel
	actorRewardsInfoRecordExecutor.SubActorFuelTx(actor, feed_attributes.Fuel(fuelCost))

	// Update Actor Reputation For the postHash
	actorVotesCountersRecord := actor_votes_counters_record_config.ActorVotesCountersRecord{
		Actor:            actor,
		PostHash:         postHash,
		LatestReputation: actorRewardsInfo.Reputation,
		LatestVoteType:   voteType,
	}
	upsertedPostReputationsRecord := actorVotesCountersRecordExecutor.UpsertActorVotesCountersRecordTx(
		&actorVotesCountersRecord)

	// Record current vote
	postVotesRecord.DeltaFuel = int64(fuelCost.Neg())
	postVotesRecord.DeltaReputation = 0
	postVotesRecord.DeltaMilestonePoints = 0
	postVotesRecord.SignedReputation = actorRewardsInfo.Reputation.Value() * postVotesRecord.VoteType.Value()
	postVotesRecord.PostType = string(postExecutor.GetPostTypeTx(postHash))
	postVotesRecordExecutor.UpsertPostVotesRecordTx(postVotesRecord)

	newPostVotesCountersRecord := post_votes_counters_record_config.PostVotesCountersRecord{
		PostHash:              postHash,
		LatestVoteType:        voteType,
		LatestActorReputation: actorRewardsInfo.Reputation,
	}
	upsertPostVotesCountersRecord := postVotesCountersRecordExecutor.UpsertPostVotesCountersRecordTx(
		&newPostVotesCountersRecord)

	voteInfo.RewardsInfo = actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)

	voteInfo.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
		UpVoteCount:    upsertPostVotesCountersRecord.UpVoteCount,
		DownVoteCount:  upsertPostVotesCountersRecord.DownVoteCount,
		TotalVoteCount: upsertPostVotesCountersRecord.TotalVoteCount,
	}

	voteInfo.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
		UpVoteCount:    upsertedPostReputationsRecord.UpVoteCount,
		DownVoteCount:  upsertedPostReputationsRecord.DownVoteCount,
		TotalVoteCount: upsertedPostReputationsRecord.TotalVoteCount,
	}

	if totalReputationsForPostHashWithSameVoteType > 0 {
		// Distribute Rewards
		for _, actorAddress := range actorList {
			awardedActorReputation :=
				actorVotesCountersRecordExecutor.GetReputationByPostHashAndActorWithLatestVoteTypeAndTimeCutOffTx(
					postVotesRecord.PostHash,
					actorAddress,
					voteType,
					cutOffTimeStamp)
			rewards := int64(float64(fuelCost) * float64(awardedActorReputation) /
				float64(totalReputationsForPostHashWithSameVoteType))

			log.Printf("rewards %+v for actorAddress %s\n", rewards, actorAddress)
			actorRewardsInfoRecordExecutor.AddActorReputationTx(actorAddress, feed_attributes.Reputation(rewards))
			actorRewardsInfoRecordExecutor.AddActorMilestonePointsFromVotesTx(
				actorAddress, feed_attributes.MilestonePoint(rewards))
			postVotesRecordExecutor.AddPostVoteDeltaRewardsInfoTx(
				actorAddress, postHash, voteType, 0, int64(rewards), int64(rewards))
		}
	}

	return &voteInfo
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
