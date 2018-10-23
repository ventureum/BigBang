package lambda_tcr_table_creator_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/TCR/project_config"
  "BigBang/internal/platform/postgres_config/TCR/proxy_config"
  "BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
  "BigBang/internal/platform/postgres_config/TCR/objective_config"
  "BigBang/internal/platform/postgres_config/TCR/milestone_config"
  "BigBang/internal/platform/postgres_config/TCR/actor_delegate_votes_account_config"
  "BigBang/internal/platform/postgres_config/TCR/principal_proxy_votes_config"
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
  objectiveExecutor := objective_config.ObjectiveExecutor{*postgresBigBangClient}
  milestoneExecutor := milestone_config.MilestoneExecutor{*postgresBigBangClient}
  proxyExecutor := proxy_config.ProxyExecutor{*postgresBigBangClient}
  ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
  actorDelegateVotesAccountExecutor := actor_delegate_votes_account_config.ActorDelegateVotesAccountExecutor{*postgresBigBangClient}
  principalProxyVotesExecutor := principal_proxy_votes_config.PrincipalProxyVotesExecutor{*postgresBigBangClient}

  principalProxyVotesExecutor.DeletePrincipalProxyVotesTable()
  actorDelegateVotesAccountExecutor.DeleteActorRatingVoteAccountTable()
  objectiveExecutor.DeleteObjectiveTable()
  milestoneExecutor.DeleteMilestoneTable()
  projectExecutor.DeleteProjectTable()
  proxyExecutor.DeleteProxyTable()
  ratingVoteExecutor.DeleteRatingVoteTable()


  projectExecutor.CreateProjectTable()
  milestoneExecutor.CreateMilestoneTable()
  objectiveExecutor.CreateObjectiveTable()
  proxyExecutor.CreateProxyTable()
  ratingVoteExecutor.CreateRatingVoteTable()
  actorDelegateVotesAccountExecutor.CreateActorRatingVoteAccountTable()
  principalProxyVotesExecutor.CreatePrincipalProxyVotesTable()

  postgresBigBangClient.Commit()
  response.Ok = true
}

func Handler() (response Response, err error) {
  response.Ok = false
  ProcessRequest(&response)
  return response, nil
}
