package lambda_get_batch_rating_vote_list_config

import (
	"BigBang/cmd/lambda/common/auth"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/platform/postgres_config/TCR/objective_config"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"log"
)

type Request struct {
	PrincipalId string         `json:"principalId,required"`
	Body        RequestContent `json:"body,required"`
}

type RequestContent struct {
	ObjectiveVotesInfoKeyList []tcr_attributes.ObjectiveVotesInfoKey `json:"objectiveVotesInfoKeyList,required"`
}

type Response struct {
	ObjectiveVotesInfoList *[]tcr_attributes.ObjectiveVotesInfo `json:"objectiveVotesInfoList,omitempty"`
	Ok                     bool                                 `json:"ok"`
	Message                *error_config.ErrorInfo              `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	defer func() {
		if errPanic := recover(); errPanic != nil { //catch
			response.ObjectiveVotesInfoList = nil
			response.Message = error_config.CreatedErrorInfoFromString(errPanic)
			postgresBigBangClient.RollBack()
		}
		postgresBigBangClient.Close()
	}()
	postgresBigBangClient.Begin()
	auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

	objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

	objectiveVotesInfoKeyList := request.Body.ObjectiveVotesInfoKeyList

	for _, objectiveVotesInfoKey := range objectiveVotesInfoKeyList {
		projectId := objectiveVotesInfoKey.ProjectId
		milestoneId := objectiveVotesInfoKey.MilestoneId
		objectiveId := objectiveVotesInfoKey.ObjectiveId
		objectiveExecutor.VerifyObjectiveRecordExistingTx(projectId, milestoneId, objectiveId)
	}

	var objectiveVotesInfoList []tcr_attributes.ObjectiveVotesInfo

	for _, objectiveVotesInfoKey := range objectiveVotesInfoKeyList {
		projectId := objectiveVotesInfoKey.ProjectId
		milestoneId := objectiveVotesInfoKey.MilestoneId
		objectiveId := objectiveVotesInfoKey.ObjectiveId
		ratingVotes := ratingVoteExecutor.GetRatingVotesByIDsTx(
			projectId, milestoneId, objectiveId)
		objectiveVotesInfo := &tcr_attributes.ObjectiveVotesInfo{
			ProjectId:   projectId,
			MilestoneId: milestoneId,
			ObjectiveId: objectiveId,
			RatingVotes: ratingVotes,
		}

		objectiveVotesInfoList = append(objectiveVotesInfoList, *objectiveVotesInfo)

		log.Printf("ObjectiveVotesInfo is loaded for ProjectId %s, MilestoneId %d, and ObjectiveId %d\n",
			projectId, milestoneId, objectiveId)
	}

	response.ObjectiveVotesInfoList = &objectiveVotesInfoList

	postgresBigBangClient.Commit()
	response.Ok = true
}

func Handler(request Request) (response Response, err error) {
	response.Ok = false
	ProcessRequest(request, &response)
	return response, nil
}
