package lambda_feed_redeem_milestone_points_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_milestone_points_redeem_history_record_config"
	"BigBang/internal/platform/postgres_config/feed/redeem_block_info_record_config"
)

// Request only used for test
type Request struct {
	IncreasedRedeemBlockNum int64 `json:"increasedRedeemBlockNum,required"`
}

type Response struct {
	Ok      bool                    `json:"ok"`
	Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()
	redeemBlockInfoRecordExecutor := redeem_block_info_record_config.RedeemBlockInfoRecordExecutor{*postgresBigBangClient}
	actorMilestonePointsRedeemHistoryRecordExecutor := actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecordExecutor{*postgresBigBangClient}

	currentRedeemBlock := feed_attributes.MoveToNextNRedeemBlock(request.IncreasedRedeemBlockNum)
	redeemBlockInfoRecordExecutor.UpdateTotalEnrolledMilestonePointsForRedeemBlockInfoRecordTx(currentRedeemBlock + 1)

	existingInActorMilestonePointsRedeemHistoryRecord := actorMilestonePointsRedeemHistoryRecordExecutor.VerifyRedeemBlockExistingTx(currentRedeemBlock)

	if !existingInActorMilestonePointsRedeemHistoryRecord {
		actorMilestonePointsRedeemHistoryRecordExecutor.UpsertBatchActorMilestonePointsRedeemHistoryRecordByRedeemBlockTx(currentRedeemBlock)
		redeemBlockInfoRecordExecutor.InitRedeemBlockInfoTx(currentRedeemBlock + 1)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
