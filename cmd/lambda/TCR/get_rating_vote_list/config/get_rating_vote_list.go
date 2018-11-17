package lambda_get_rating_vote_list_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
)


type Request struct {
  ProjectId     string         `json:"projectId,required"`
  MilestoneId   int64          `json:"milestoneId,required"`
  ObjectiveId   int64          `json:"objectiveId,required"`
  Limit  int64  `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}

type ResponseData struct {
  ObjectiveVotesInfo *tcr_attributes.ObjectiveVotesInfo `json:"objectiveVotesInfo,omitempty"`
  NextCursor string                              `json:"nextCursor,omitempty"`
}

type Response struct {
  ResponseData *ResponseData`json:"responseData,omitempty"`
  Ok bool                                        `json:"ok"`
  Message *error_config.ErrorInfo                `json:"message,omitempty"`
}


func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.ResponseData = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId
  objectiveId := request.ObjectiveId
  limit := request.Limit

  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
  objectiveExecutor.VerifyObjectiveRecordExisting(projectId, milestoneId, objectiveId)

  cursorStr := request.Cursor
  var cursor string
  if cursorStr != "" {
    cursor = utils.Base64DecodeToString(cursorStr)
  }

  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

  ratingVoteRecords := ratingVoteExecutor.GetRatingVoteRecordsByCursor(
    projectId, milestoneId, objectiveId, cursor, limit + 1)

  response.ResponseData = &ResponseData{
    NextCursor: "",
    ObjectiveVotesInfo: &tcr_attributes.ObjectiveVotesInfo{
      ProjectId: projectId,
      MilestoneId: milestoneId,
      ObjectiveId: objectiveId,
    },
  }

  var ratingVotes []tcr_attributes.RatingVote
  for index, ratingVoteRecord := range *ratingVoteRecords {
    if index < int(limit) {
      ratingVote := tcr_attributes.RatingVote{
        Voter: ratingVoteRecord.Voter,
        Rating: ratingVoteRecord.Rating,
        Weight: ratingVoteRecord.Weight,
        BlockTimestamp: ratingVoteRecord.BlockTimestamp,
      }
      ratingVotes = append(ratingVotes, ratingVote)
    } else {
      response.ResponseData.NextCursor = ratingVoteRecord.EncodeID()
    }
  }


  response.ResponseData.ObjectiveVotesInfo.RatingVotes = &ratingVotes

  if cursorStr == "" {
    log.Printf("ObjectiveVotesInfo is loaded for first query with ProjectId %s, MilestoneId %d, ObjectiveId %d and limit %d\n",
      projectId, milestoneId, objectiveId, limit)
  } else {
    log.Printf("ObjectiveVotesInfo is loaded for query with ProjectId %s, MilestoneId %d, ObjectiveId %d, cursor %s and limit %d\n",
      projectId, milestoneId, objectiveId, cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
