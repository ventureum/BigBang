package lambda_get_rating_vote_list_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
)


type Request struct {
  ProjectId     string         `json:"projectId,required"`
  MilestoneId   int64          `json:"milestoneId,required"`
  ObjectiveId   int64          `json:"objectiveId,required"`
  Limit  int64  `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}


type Response struct {
  ObjVoteInfo *tcr_attributes.ObjVoteInfo `json:"objVoteInfo,omitempty"`
  NextCursor string `json:"nextCursor,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}


func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.ObjVoteInfo = nil
      response.NextCursor = ""
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  projectId := request.ProjectId
  milestoneId := request.MilestoneId
  objectiveId := request.ObjectiveId
  limit := request.Limit
  cursorStr := request.Cursor

  var cursor int64
  if cursorStr != "" {
    cursor = utils.Base64DecodeToInt64(cursorStr)
  }

  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

  ratingVoteRecords := ratingVoteExecutor.GetRatingVoteRecordsByCursor(
    projectId, milestoneId, objectiveId, cursor, limit + 1)

  response.NextCursor = ""
  response.ObjVoteInfo = &tcr_attributes.ObjVoteInfo{
    ProjectId: projectId,
    MilestoneId: milestoneId,
    ObjectiveId: objectiveId,
  }

  var ratingVotes []tcr_attributes.RatingVote
  for index, ratingVoteRecord := range *ratingVoteRecords {
    if index < int(limit) {
      ratingVote := tcr_attributes.RatingVote{
        Voter: ratingVoteRecord.Voter,
        Rating: ratingVoteRecord.Rating,
        Weight: ratingVoteRecord.Weight,
        VotedAt: ratingVoteRecord.CreatedAt,
      }
      ratingVotes = append(ratingVotes, ratingVote)
    } else {
      response.NextCursor = utils.Base64EncodeInt64(ratingVoteRecord.ID)
    }
  }


  response.ObjVoteInfo.RatingVotes = &ratingVotes

  if cursorStr == "" {
    log.Printf("ObjVoteInfo is loaded for first query with ProjectId %s, MilestoneId %d, ObjectiveId %d and limit %d\n",
      projectId, milestoneId, objectiveId, limit)
  } else {
    log.Printf("ObjVoteInfo is loaded for query with ProjectId %s, MilestoneId %d, ObjectiveId %d, cursor %s and limit %d\n",
      projectId, milestoneId, objectiveId, cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
