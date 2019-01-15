package lambda_get_next_redeem_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
	"BigBang/internal/platform/postgres_config/feed/milestone_points_redeem_request_record_config"
	"BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
	"log"
	"time"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor string `json:"actor,required"`
}

type ResponseContent struct {
	Actor                   string                           `json:"actor,required"`
	TargetedMilestonePoints feed_attributes.MilestonePoint   `json:"targetedMilestonePoints,required"`
	ActualMilestonePoints   feed_attributes.MilestonePoint   `json:"actualMilestonePoints,required"`
	EstimatedTokens         int64                            `json:"estimatedTokens,required"`
	SubmittedAt             time.Time                        `json:"submittedAt,required"`
	RedeemBlockInfo         *feed_attributes.RedeemBlockInfo `json:"redeemBlockInfo,omitempty"`
}

type Response struct {
	NextRedeem *ResponseContent        `json:"nextRedeem,omitempty"`
	Ok         bool                    `json:"ok"`
	Message    *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.NextRedeem = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)

	milestonePointsRedeemRequestRecordExecutor :=
		milestone_points_redeem_request_record_config.MilestonePointsRedeemRequestRecordExecutor{
			*postgresBigBangClient}
	actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
		*postgresBigBangClient}
	redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{
		*postgresBigBangClient}

	nextRedeemBlock := feed_attributes.MoveToNextNRedeemBlock(1)
	milestonePointsRedeemRequestRecordExecutor.VerifyMilestonePointsRedeemRequestExistingTx(actor)

	milestonePointsRedeemRequest := milestonePointsRedeemRequestRecordExecutor.GetMilestonePointsRedeemRequestTx(actor)

	redeemBlockInfo := redeemBlockInfoRecordExecutor.GetRedeemBlockInfoTx(nextRedeemBlock)
	actorRewardsInfoRecord := actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)

	milestonePointsToRedeem := actorRewardsInfoRecord.MilestonePoints
	targetedMilestonePoints := milestonePointsRedeemRequest.TargetedMilestonePoints

	if milestonePointsToRedeem > milestonePointsRedeemRequest.TargetedMilestonePoints {
		milestonePointsToRedeem = targetedMilestonePoints
	}

	var estimatedTokens int64

	if redeemBlockInfo.TotalEnrolledMilestonePoints > 0 {
		estimatedTokens = redeemBlockInfo.TokenPool * int64(milestonePointsToRedeem) / int64(redeemBlockInfo.TotalEnrolledMilestonePoints)
	}

	response.NextRedeem = &ResponseContent{
		Actor:                   actor,
		TargetedMilestonePoints: targetedMilestonePoints,
		ActualMilestonePoints:   actorRewardsInfoRecord.MilestonePoints,
		EstimatedTokens:         estimatedTokens,
		SubmittedAt:             milestonePointsRedeemRequest.SubmittedAt,
		RedeemBlockInfo:         redeemBlockInfo,
	}

	log.Printf("Sucessfully loaded content for get_next_redeem for actor %s\n", actor)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
