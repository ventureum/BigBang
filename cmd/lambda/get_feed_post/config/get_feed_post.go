package config

import (
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/post_votes_counters_record_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/platform/postgres_config/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/actor_votes_counters_record_config"
)


type Request struct {
  PostHash string `json:"postHash,required"`
  Requestor string `json:"requestor,omitempty"`
}

type ResponseContent struct {
  Actor string `json:"actor"`
  Username string `json:"username,required"`
  PhotoUrl string `json:"photoUrl,required"`
  BoardId string `json:"boardId"`
  ParentHash string `json:"parentHash"`
  PostHash string `json:"postHash"`
  PostType string `json:"postType"`
  Content *feed_attributes.Content `json:"content"`
  DeltaFuel int64 `json:"deltaFuel"`
  DeltaReputation int64 `json:"deltaReputation"`
  DeltaMilestonePoints int64 `json:"deltaMilestonePoints"`
  WithdrawableMPs int64 `json:"withdrawableMPs"`
  RepliesLength int64 `json:"repliesLength"`
}

type Response struct {
  Post *ResponseContent `json:"post,omitempty"`
  PostVoteCountInfo *feed_attributes.VoteCountInfo `json:"postVoteCountInfo,omitempty"`
  RequestorVoteCountInfo *feed_attributes.VoteCountInfo `json:"requestorVoteCountInfo,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func PostRecordResultToResponseContent(result *post_config.PostRecordResult) *ResponseContent {
  return &ResponseContent{
    Actor: result.Actor,
    BoardId: result.BoardId,
    ParentHash: result.ParentHash,
    PostHash: result.PostHash,
    PostType: result.PostType,
    Content: result.Content,
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Post = nil
      response.PostVoteCountInfo = nil
      response.RequestorVoteCountInfo = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresFeedClient.Close()
  }()

  postHash := request.PostHash
  requestor := request.Requestor

  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}
  postVotesCounterRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresFeedClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresFeedClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresFeedClient}

  postExecutor.VerifyPostRecordExisting(postHash)
  if requestor != "" {
    actorProfileRecordExecutor.VerifyActorExisting(requestor)
    actorRewardsInfoRecordExecutor.VerifyActorExisting(requestor)
  }

  postRecordResult := postExecutor.GetPostRecord(postHash).ToPostRecordResult()
  response.Post = PostRecordResultToResponseContent(postRecordResult)
  actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecord(postRecordResult.Actor)
  response.Post.Username = actorProfileRecord.Username
  response.Post.PhotoUrl = actorProfileRecord.PhotoUrl
  response.Post.RepliesLength = postRepliesRecordExecutor.GetPostRepliesRecordCount(postHash)
  postRewardsRecord := postRewardsRecordExecutor.GetPostRewardsRecordByPostHash(postHash)
  response.Post.DeltaFuel = postRewardsRecord.DeltaFuel
  response.Post.DeltaReputation = postRewardsRecord.DeltaReputation
  response.Post.DeltaMilestonePoints = postRewardsRecord.DeltaMilestonePoints
  response.Post.WithdrawableMPs = postRewardsRecord.WithdrawableMPs


  log.Printf("Post Content is loaded for postHash %s\n", postHash)

  postVotesCounterRecord := postVotesCounterRecordExecutor.GetPostVotesCountersRecordByPostHash(postHash)
  response.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
    DownVoteCount:  postVotesCounterRecord.DownVoteCount,
    UpVoteCount:    postVotesCounterRecord.UpVoteCount,
    TotalVoteCount: postVotesCounterRecord.TotalVoteCount,
  }

  log.Printf("PostVoteInfo is loaded for postHash %s\n", postHash)

  if requestor != "" {
    actorVotesCountersRecord := actorVotesCountersRecordExecutor.GetActorVotesCountersRecordByPostHashAndActor(postHash, requestor)
    response.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
      DownVoteCount:  actorVotesCountersRecord.DownVoteCount,
      UpVoteCount:    actorVotesCountersRecord.UpVoteCount,
      TotalVoteCount: actorVotesCountersRecord.TotalVoteCount,
    }
    log.Printf("RequestorVoteInfo is loaded for postHash %s\n", postHash)
  }

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
