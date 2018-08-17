package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
)


type Request struct {
  PostHash string `json:"postHash,required"`
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
  Post ResponseContent `json:"post,omitempty"`
  Ok bool `json:"ok"`
  Message string `json:"message,omitempty"`
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
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      response.Message = errStr.(string)
    }
  }()

  postHash := request.PostHash
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer postgresFeedClient.Close()
  postgresFeedClient.Begin()

  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}

  postRecordResult := postExecutor.GetPostRecordTx(postHash).ToPostRecordResult()

  response.Post = *PostRecordResultToResponseContent(postRecordResult)
  response.Post.RepliesLength = postRepliesRecordExecutor.GetPostRepliesRecordCountTx(postHash)
  response.Post.Rewards = postRewardsRecordExecutor.GetPostRewardsTx(postHash).Value()

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
