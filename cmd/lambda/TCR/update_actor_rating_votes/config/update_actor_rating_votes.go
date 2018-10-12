package lambda_update_actor_rating_votes_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/actor_rating_vote_account_config"
)


type Request struct {
  Actor   string  `json:"actor,required"`
  ProjectId string `json:"projectId,required"`
  AvailableVotes int64  `json:"availableVotes,required"`
  ReceivedVotes int64 `json:"receivedVotes,required"`
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
  actorRatingVoteAccountExecutor := actor_rating_vote_account_config.ActorRatingVoteAccountExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(request.Actor)

  actorRatingVoteAccountExecutor.UpsertActorRatingVoteAccountRecordTx(&actor_rating_vote_account_config.ActorRatingVoteAccountRecord{
    Actor: request.Actor,
    ProjectId: request.ProjectId,
    AvailableRatingVotes: request.AvailableVotes,
    ReceivedRatingVotes: request.ReceivedVotes,
  })

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
