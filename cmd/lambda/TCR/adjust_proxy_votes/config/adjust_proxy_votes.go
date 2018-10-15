package lambda_adjust_proxy_votes_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
)


type Request struct {
  Actor   string  `json:"actor,required"`
  ProjectId string `json:"projectId,required"`
  Proxy string `json:"proxy,required"`
  BlockTimestamp  int64        `json:"blockTimestamp,required"`
  Votes int64  `json:"votes,required"`
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
  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(request.Actor)
  actorProfileRecordExecutor.VerifyActorExistingTx(request.Proxy)
  projectExecutor.VerifyProjectRecordExistingTx(request.ProjectId)

  if request.Votes > 0 {
    principalProxyVotesRecord := &principal_proxy_votes_config.PrincipalProxyVotesRecord{
      Actor: request.Actor,
      ProjectId: request.ProjectId,
      Proxy: request.Proxy,
      BlockTimestamp: request.BlockTimestamp,
      Votes: request.Votes,
    }
    principalProxyVotesRecord.GenerateID()
    principalProxyVotesExecutor.UpsertPrincipalProxyVotesRecordTx(principalProxyVotesRecord)
  } else {
    principalProxyVotesExecutor.DeletePrincipalProxyVotesRecordByIDsTx(request.Actor, request.ProjectId, request.Proxy)
  }

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
