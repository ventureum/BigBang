package lambda_update_batch_available_delegate_votes_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
)

type RequestContent struct {
  Actor                  string `json:"actor,required"`
  ProjectId              string `json:"projectId,required"`
  AvailableDelegateVotes int64  `json:"availableDelegateVotes,required"`
}

type Request struct {
  RequestList []RequestContent  `json:"requestList,required"`
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


  requestList := request.RequestList


  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}


  for _ , singleRequest := range requestList {
    actorProfileRecordExecutor.VerifyActorExistingTx(singleRequest.Actor)
    projectExecutor.VerifyProjectRecordExistingTx(singleRequest.ProjectId)
  }


  for _ , singleRequest := range requestList {
    actor := singleRequest.Actor
    projectId := singleRequest.ProjectId
    availableDelegateVotes := singleRequest.AvailableDelegateVotes
    existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExistingTx(actor, projectId)

    if !existing {
      actorDelegateVotesAccountExecutor.UpsertActorDelegateVotesAccountRecordTx(&actor_delegate_votes_account_config.ActorDelegateVotesAccountRecord{
        Actor: actor,
        ProjectId: projectId,
        AvailableDelegateVotes: availableDelegateVotes,
        ReceivedDelegateVotes: 0,
      })
    } else {
      actorDelegateVotesAccountExecutor.UpdateAvailableDelegateVotesTx(actor, projectId, availableDelegateVotes)
    }

  }

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
