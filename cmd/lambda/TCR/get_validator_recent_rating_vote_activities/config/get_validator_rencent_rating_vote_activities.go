package lambda_get_validator_recent_rating_vote_activities_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/internal/platform/postgres_config/client_config"
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
	RatingVoteActivities *[]tcr_attributes.RatingVoteActivity `json:"ratingVoteActivities,omitempty"`
	NextCursor           string                               `json:"nextCursor,omitempty"`
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
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	actor := request.Body.Actor
	limit := request.Body.Limit

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{
		*postgresBigBangClient}
	actorProfileRecordExecutor.VerifyActorExistingTx(actor)

	cursorStr := request.Body.Cursor
	var cursor string
	if cursorStr != "" {
		cursor = utils.Base64DecodeToString(cursorStr)
	}

	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

	ratingVoteActivities := ratingVoteExecutor.GetRatingVoteActivitiesForActorByCursorTx(actor, cursor, limit+1)

	response.ResponseData = &ResponseData{
		NextCursor:           "",
		RatingVoteActivities: nil,
	}

	if ratingVoteActivities != nil {
		ratingVoteActivitiesLen := int64(len(*ratingVoteActivities))
		if ratingVoteActivitiesLen <= limit {
			var newRatingVoteActivities = (*ratingVoteActivities)[0:ratingVoteActivitiesLen]
			response.ResponseData.RatingVoteActivities = &newRatingVoteActivities
		} else {
			var newRatingVoteActivities = (*ratingVoteActivities)[0:limit]
			response.ResponseData.RatingVoteActivities = &newRatingVoteActivities
			lastActivity := (*ratingVoteActivities)[limit]
			response.ResponseData.NextCursor = rating_vote_config.GenerateEncodedRatingVoteRecordID(
				lastActivity.ProjectId,
				lastActivity.MilestoneId,
				lastActivity.ObjectiveId,
				actor,
				lastActivity.BlockTimestamp,
			)
		}
	}

	if cursorStr == "" {
		log.Printf("RatingVoteActivities is loaded for first query with actor %s and limit %d\n",
			actor, limit)
	} else {
		log.Printf("RatingVoteActivities is loaded for query with actor %s, cursor %s and limit %d\n",
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
