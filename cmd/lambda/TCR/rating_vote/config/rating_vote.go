package lambda_rating_vote_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/milestone_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/TCR/project_config"
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
	ProjectId      string `json:"projectId,required"`
	MilestoneId    int64  `json:"milestoneId,required"`
	ObjectiveId    int64  `json:"objectiveId,required"`
	Voter          string `json:"voter,required"`
	BlockTimestamp int64  `json:"blockTimestamp,required"`
	Rating         int64  `json:"rating,required"`
	Weight         int64  `json:"weight,required"`
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
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	projectId := request.Body.ProjectId
	milestoneId := request.Body.MilestoneId
	objectiveId := request.Body.ObjectiveId
	voter := request.Body.Voter
	rating := request.Body.Rating
	weight := request.Body.Weight

	actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
	milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
	projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}

	actorProfileRecordExecutor.VerifyActorExistingTx(voter)
	existing := ratingVoteExecutor.VerifyRatingVoteRecordExistingTx(projectId, milestoneId, objectiveId, voter)
	if existing {
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.RatingVoteExceedingLimitedVotingTimes,
			ErrorData: map[string]interface{}{
				"objectiveId": objectiveId,
				"milestoneId": milestoneId,
				"projectId":   projectId,
				"voter":       voter,
			},
			ErrorLocation: error_config.RatingVoteRecordLocation,
		}
		log.Printf("Rating Vote Exceeds Limited Voting Times for projectId %s, milestoneId %d, objectiveId %d by voter %s",
			projectId, milestoneId, objectiveId, voter)
		log.Panicln(errorInfo.Marshal())
	}

	objectiveExecutor.VerifyObjectiveRecordExistingTx(projectId, milestoneId, objectiveId)
	ratingVoteRecord := rating_vote_config.RatingVoteRecord{
		ProjectId:      projectId,
		MilestoneId:    milestoneId,
		ObjectiveId:    objectiveId,
		Voter:          voter,
		BlockTimestamp: request.Body.BlockTimestamp,
		Rating:         rating,
		Weight:         weight,
	}

	ratingVoteRecord.GenerateID()
	ratingVoteExecutor.UpsertRatingVoteRecordTx(&ratingVoteRecord)
	objectiveExecutor.AddRatingAndWeightForObjectiveTx(projectId, milestoneId, objectiveId, rating, weight)
	milestoneExecutor.AddRatingAndWeightForMilestoneTx(projectId, milestoneId, rating, weight)
	projectExecutor.AddRatingAndWeightForProjectTx(projectId, rating, weight)

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
