package config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/platform/postgres_config/TCR/proxy_config"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
)

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  postgresBigBangClient.Begin()
  postgresBigBangClient.SetIdleInTransactionSessionTimeout(60000)

  projectExecutor := project_config.ProjectExecutor{*postgresBigBangClient}
  proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}

  projectExecutor.DeleteProjectTable()
  proxyExecutor.DeleteProxyTable()
  ratingVoteExecutor.DeleteRatingVoteTable()

  projectExecutor.CreateProjectTable()
  proxyExecutor.CreateProxyTable()
  ratingVoteExecutor.CreateRatingVoteTable()

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}
