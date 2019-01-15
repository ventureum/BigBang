package lambda_get_redeem_history_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/internal/platform/postgres_config/feed/actor_milestone_points_redeem_history_record_config"
	"BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	Actor  string `json:"actor,required"`
	Limit  int64  `json:"limit,required"`
	Cursor string `json:"cursor,omitempty"`
}

type ResponseData struct {
	Redeems    *[]feed_attributes.MilestonePointsRedeemHistory `json:"redeems,omitempty"`
	NextCursor string                                          `json:"nextCursor,omitempty"`
}

type Response struct {
	ResponseData *ResponseData           `json:"responseData,omitempty"`
	Ok           bool                    `json:"ok"`
	Message      *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.ResponseData = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()

	postgresBigBangClient.Begin()

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

	actor := request.Body.Actor
	auth.AuthProcess(request.PrincipalId, actor, postgresBigBangClient)
	limit := request.Body.Limit
	cursorStr := request.Body.Cursor

	actorProfileRecordExecutor.VerifyActorExistingTx(actor)

	var cursor string
	if cursorStr != "" {
		cursor = utils.Base64DecodeToString(cursorStr)
	}

	actorMilestonePointsRedeemHistoryRecordExecutor :=
		actor_milestone_points_redeem_history_record_config.ActorMilestonePointsRedeemHistoryRecordExecutor{
			*postgresBigBangClient}

	actorMilestonePointsRedeemHistory :=
		actorMilestonePointsRedeemHistoryRecordExecutor.GetActorMilestonePointsRedeemHistoryByCursorTx(
			actor, cursor, limit+1)

	response.ResponseData = &ResponseData{
		NextCursor: "",
		Redeems:    nil,
	}

	var redeems []feed_attributes.MilestonePointsRedeemHistory
	for index, redeem := range *actorMilestonePointsRedeemHistory {
		if index < int(limit) {
			redeems = append(redeems, redeem)
		} else {
			response.ResponseData.NextCursor = redeem.GenerateRecordID()
		}
	}

	response.ResponseData.Redeems = &redeems

	if cursorStr == "" {
		log.Printf("ActorMilestonePointsRedeemHistory is loaded for first query with actor %s and limit %d\n",
			actor, limit)
	} else {
		log.Printf("ActorMilestonePointsRedeemHistory is loaded for query with actor %s, cursor %s and limit %d\n",
			actor, cursorStr, limit)
	}

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
