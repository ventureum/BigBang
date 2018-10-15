package lambda_get_proxy_votes_list_by_principal_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/pkg/utils"
  "BigBang/internal/app/tcr_attributes"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
)


type Request struct {
  Actor string  `json:"actor,required"`
  ProjectId     string    `json:"projectId,required"`
  Limit  int64  `json:"limit,required"`
  Cursor string `json:"cursor,omitempty"`
}


type Response struct {
  ProxyVotesInfo *tcr_attributes.ProxyVotesInfo  `json:"proxyVotesInfo,omitempty"`
  NextCursor string                              `json:"nextCursor,omitempty"`
  Ok bool                                        `json:"ok"`
  Message *error_config.ErrorInfo                `json:"message,omitempty"`
}


func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.ProxyVotesInfo = nil
      response.NextCursor = ""
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  actor := request.Actor
  projectId := request.ProjectId
  limit := request.Limit

  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExisting(actor)
  projectExecutor.VerifyProjectRecordExisting(projectId)

  principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}

  cursorStr := request.Cursor
  var cursor string
  if cursorStr != "" {
    cursor = utils.Base64DecodeToString(cursorStr)
  }

  principalProxyVotesRecordList := principalProxyVotesExecutor.GetPrincipalProxyVotesRecordListByCursor(actor, projectId, cursor, limit + 1)

  response.NextCursor = ""
  response.ProxyVotesInfo = &tcr_attributes.ProxyVotesInfo {
    Actor: actor,
    ProjectId: projectId,
  }

  var proxyVotesList []tcr_attributes.ProxyVotes
  for index, principalProxyVotesRecord := range *principalProxyVotesRecordList {
    if index < int(limit) {
      ratingVote := tcr_attributes.ProxyVotes {
        Proxy: principalProxyVotesRecord.Proxy,
        BlockTimestamp: principalProxyVotesRecord.BlockTimestamp,
        Votes: principalProxyVotesRecord.Votes,
      }
      proxyVotesList = append(proxyVotesList, ratingVote)
    } else {
      response.NextCursor = principalProxyVotesRecord.EncodeID()
    }
  }


  response.ProxyVotesInfo.ProxyVotesList = &proxyVotesList

  if cursorStr == "" {
    log.Printf("ProxyVotesInfo is loaded for first query with Actor %s, ProjectId %s and limit %d\n",
      actor, projectId, limit)
  } else {
    log.Printf("ProxyVotesInfo is loaded for query with Actor %s, ProjectId %s, cursor %s and limit %d\n",
      actor, projectId, cursorStr, limit)
  }
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
