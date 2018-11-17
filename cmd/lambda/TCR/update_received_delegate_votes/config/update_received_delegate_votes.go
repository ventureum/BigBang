package lambda_update_received_delegate_votes_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
)


type Request struct {
  Actor                  string `json:"actor,required"`
  ProjectId              string `json:"projectId,required"`
  ReceivedDelegateVotesDelta int64  `json:"receivedDelegateVotesDelta,required"`
}

type Response struct {
  Ok bool `json:"ok"`
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

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(request.Actor)
  projectExecutor.VerifyProjectRecordExistingTx(request.ProjectId)

  existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(request.Actor, request.ProjectId)

  if !existing {
    actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
      Actor: request.Actor,
      ProjectId: request.ProjectId,
      AvailableDelegateVotes: 0,
      ReceivedDelegateVotes: request.ReceivedDelegateVotesDelta,
    })
  } else {
    actorDelegateVotesAccountExecutor.UpdateReceivedDelegateVotesTx(request.Actor, request.ProjectId, request.ReceivedDelegateVotesDelta)
  }

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
