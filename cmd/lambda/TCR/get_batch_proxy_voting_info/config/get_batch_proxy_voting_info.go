package lambda_get_batch_proxy_voting_info_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
  "BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
  "BigBang/cmd/lambda/common/auth"
)

type Request struct {
  PrincipalId string `json:"principalId,required"`
  Body RequestContent `json:"body,required"`
}


type RequestContent struct {
  ProxyVotingInfoKeyList []tcr_attributes.ProxyVotingInfoKey `json:"proxyVotingInfoKeyList,required"`
}


type Response struct {
  ProxyVotingInfoList *[]tcr_attributes.ProxyVotingInfo `json:"proxyVotingInfoList,omitempty"`
  Ok              bool                            `json:"ok"`
  Message         *error_config.ErrorInfo         `json:"message,omitempty"`
}


func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.ProxyVotingInfoList  = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  auth.AuthProcess(request.PrincipalId, "", postgresBigBangClient)

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
  principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}

  proxyVotingInfoKeyList := request.Body.ProxyVotingInfoKeyList

  for _ , proxyVotingInfoKey := range proxyVotingInfoKeyList {
    actor := proxyVotingInfoKey.Actor
    projectId := proxyVotingInfoKey.ProjectId
    actorProfileRecordExecutor.VerifyActorExisting(actor)
    projectExecutor.VerifyProjectRecordExisting(projectId)
  }
  var proxyVotingInfoList []tcr_attributes.ProxyVotingInfo

  for _ , proxyVotingInfoKey := range proxyVotingInfoKeyList {
    actor := proxyVotingInfoKey.Actor
    projectId := proxyVotingInfoKey.ProjectId
    existing := actorDelegateVotesAccountExecutor.VerifyDelegateVotesAccountExisting(actor, projectId)
    var proxyVotingInfo tcr_attributes.ProxyVotingInfo
    if !existing  {
      proxyVotingInfo = tcr_attributes.ProxyVotingInfo{
        Actor:                  actor,
        ProjectId:              projectId,
        AvailableDelegateVotes: 0,
        ReceivedDelegateVotes:  0,
        ProxyVotingList:        nil,
      }

    } else {
      actorDelegateVotesAccount := actorDelegateVotesAccountExecutor.GetActorDelegateVotesAccountRecord(actor, projectId)
      proxyVotingList := principalProxyVotesExecutor.GetProxyVotingListByActorAndProjectId(actor, projectId)
      proxyVotingInfo = tcr_attributes.ProxyVotingInfo{
        Actor:                  actor,
        ProjectId:              projectId,
        AvailableDelegateVotes: actorDelegateVotesAccount.AvailableDelegateVotes,
        ReceivedDelegateVotes:  actorDelegateVotesAccount.ReceivedDelegateVotes,
        ProxyVotingList:        proxyVotingList,
      }
    }
    proxyVotingInfoList = append(proxyVotingInfoList, proxyVotingInfo)
    log.Printf("ProxyVotingInfo is loaded for  Actor %s and ProjectId %s\n", actor, projectId)
  }

  response.ProxyVotingInfoList = &proxyVotingInfoList
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
