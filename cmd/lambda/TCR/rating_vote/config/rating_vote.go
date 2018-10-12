package lambda_rating_vote_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
)


type Request struct {
  ProjectId     string         `json:"projectId,required"`
  MilestoneId   int64          `json:"milestoneId,required"`
  ObjectiveId   int64          `json:"objId,required"`
  Voter         string         `json:"voter,required"`
  BlockTimestamp  int64        `json:"blockTimestamp,required"`
  Rating        int64          `json:"rating,required"`
  Weight        int64          `json:"weight,required"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()
  postgresBigBangClient.Begin()

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}

  actorProfileRecordExecutor.VerifyActorExistingTx(request.Voter)
  objectiveExecutor.VerifyObjectiveRecordExistingTx(request.ProjectId, request.MilestoneId, request.ObjectiveId)
  ratingVoteRecord := rating_vote_config.RatingVoteRecord{
    ProjectId: request.ProjectId,
    MilestoneId: request.MilestoneId,
    ObjectiveId: request.ObjectiveId,
    Voter: request.Voter,
    Rating: request.Rating,
    Weight: request.Weight,
  }

  ratingVoteRecord.GenerateID()
  ratingVoteExecutor.UpsertRatingVoteRecordTx(&ratingVoteRecord)

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
