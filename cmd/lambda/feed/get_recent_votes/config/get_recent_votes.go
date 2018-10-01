package config

import (
  "BigBang/internal/platform/postgres_config/feed/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  Limit int64 `json:"limit,omitempty"`
}

type Response struct {
  RecentVotes *[]post_votes_record_config.PostVotesRecord `json:"recentVotes,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.RecentVotes = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresFeedClient.Close()
  }()

  actor := request.Actor
  limit := request.Limit

  if limit == 0 {
    limit = 20
  }

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresFeedClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}

  actorProfileRecordExecutor.VerifyActorExisting(actor)
  actorRewardsInfoRecordExecutor.VerifyActorExisting(actor)

  response.RecentVotes = postVotesRecordExecutor.GetRecentPostVotesRecordsByActor(actor, limit)

  log.Printf("RecentPostVotesRecords is loaded for actor %s\n", actor)

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
