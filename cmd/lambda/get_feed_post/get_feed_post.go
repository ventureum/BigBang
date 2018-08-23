package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/post_reputations_record_config"
  "BigBang/internal/pkg/error_config"
  "log"
)


type Request struct {
  PostHash string `json:"postHash,required"`
  Requestor string `json:"requestor,omitempty"`
}

type ResponseContent struct {
  Actor string `json:"actor"`
  BoardId string `json:"boardId"`
  ParentHash string `json:"parentHash"`
  PostHash string `json:"postHash"`
  PostType string `json:"postType"`
  Content feed_attributes.Content `json:"content"`
  Rewards int64 `json:"rewards"`
  RepliesLength int64 `json:"repliesLength"`
}

type Response struct {
  Post *ResponseContent `json:"post,omitempty"`
  PostVoteInfo *feed_attributes.VoteInfo `json:"postVoteInfo,omitempty"`
  RequestorVoteInfo *feed_attributes.VoteInfo `json:"requestorVoteInfo,omitempty"`
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
    Content: *result.Content,
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Post = nil
      response.PostVoteInfo = nil
      response.RequestorVoteInfo = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()

  postHash := request.PostHash
  requestor := request.Requestor

  postgresFeedClient.Begin()

  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}
  postVotesCounterRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresFeedClient}
  postReputationsRecordExecutor := post_reputations_record_config.PostReputationsRecordExecutor{*postgresFeedClient}

  postRecordResult := postExecutor.GetPostRecordTx(postHash).ToPostRecordResult()
  response.Post = PostRecordResultToResponseContent(postRecordResult)
  response.Post.RepliesLength = postRepliesRecordExecutor.GetPostRepliesRecordCountTx(postHash)
  response.Post.Rewards = postRewardsRecordExecutor.GetPostRewardsTx(postHash).Value()

  log.Printf("Post Content is loaded for postHash %s\n", postHash)

  postVotesCounterRecord := postVotesCounterRecordExecutor.GetPostVotesCountersRecordByPostHashTx(postHash)
  response.PostVoteInfo = &feed_attributes.VoteInfo{
    DownVoteCount:  postVotesCounterRecord.DownVoteCount,
    UpVoteCount:    postVotesCounterRecord.UpVoteCount,
    TotalVoteCount: postVotesCounterRecord.TotalVoteCount,
  }

  log.Printf("PostVoteInfo is loaded for postHash %s\n", postHash)

  if requestor != "" {
    postReputationsRecord := postReputationsRecordExecutor.GetPostReputationsRecordByPostHashAndActorTx(postHash, requestor)
    response.RequestorVoteInfo = &feed_attributes.VoteInfo{
      DownVoteCount:  postReputationsRecord.DownVoteCount,
      UpVoteCount:    postReputationsRecord.UpVoteCount,
      TotalVoteCount: postReputationsRecord.TotalVoteCount,
    }
    log.Printf("RequestorVoteInfo is loaded for postHash %s\n", postHash)
  }

  postgresFeedClient.Commit()
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}


func main() {
  // TODO(david.shao): remove example when deployed to production
  //request := Request{
  // PostHash: "0x009",
  //}
  //response, _ := Handler(request)
  //fmt.Printf("%+v", response)

  lambda.Start(Handler)
}
