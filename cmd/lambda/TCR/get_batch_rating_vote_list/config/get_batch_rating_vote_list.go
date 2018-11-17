package lambda_get_batch_rating_vote_list_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
)


type Request struct {
  ObjectiveVotesInfoKeyList []tcr_attributes.ObjectiveVotesInfoKey `json:"objectiveVotesInfoKeyList,required"`
}


type Response struct {
  ObjectiveVotesInfoList *[]tcr_attributes.ObjectiveVotesInfo `json:"objectiveVotesInfoList,omitempty"`
  Ok bool                                        `json:"ok"`
  Message *error_config.ErrorInfo                `json:"message,omitempty"`
}


func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.ObjectiveVotesInfoList = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

  objectiveVotesInfoKeyList := request.ObjectiveVotesInfoKeyList

  for _ , objectiveVotesInfoKey := range objectiveVotesInfoKeyList {
    projectId := objectiveVotesInfoKey.ProjectId
    milestoneId := objectiveVotesInfoKey.MilestoneId
    objectiveId := objectiveVotesInfoKey.ObjectiveId
    objectiveExecutor.VerifyObjectiveRecordExisting(projectId, milestoneId, objectiveId)
  }

  var objectiveVotesInfoList []tcr_attributes.ObjectiveVotesInfo

  for _ , objectiveVotesInfoKey := range objectiveVotesInfoKeyList {
    projectId := objectiveVotesInfoKey.ProjectId
    milestoneId := objectiveVotesInfoKey.MilestoneId
    objectiveId := objectiveVotesInfoKey.ObjectiveId
    ratingVotes:= ratingVoteExecutor.GetRatingVotesByIDs(
      projectId, milestoneId, objectiveId)
    objectiveVotesInfo := &tcr_attributes.ObjectiveVotesInfo{
      ProjectId: projectId,
      MilestoneId: milestoneId,
      ObjectiveId: objectiveId,
      RatingVotes: ratingVotes,
    }

    objectiveVotesInfoList = append(objectiveVotesInfoList, *objectiveVotesInfo)

    log.Printf("ObjectiveVotesInfo is loaded for ProjectId %s, MilestoneId %d, and ObjectiveId %d\n",
      projectId, milestoneId, objectiveId)
  }

  response.ObjectiveVotesInfoList = &objectiveVotesInfoList

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
