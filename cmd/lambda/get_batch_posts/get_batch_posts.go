package main

import (
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
)


type Request struct {
  PostHashes []string `json:"postHashes,required"`
}

type ResponseContent struct {
  Actor string `json:"actor"`
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
  Posts *[] ResponseContent `json:"posts,omitempty"`
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
      response.Posts = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresFeedClient.Close()
  }()

  postHashes := request.PostHashes


  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}

  for _, postHash := range postHashes {
    postExecutor.VerifyPostRecordExisting(postHash)
  }
  posts := make([]ResponseContent, len(postHashes))
  for index, postHash := range postHashes {
    var post *ResponseContent
    postRecordResult := postExecutor.GetPostRecord(postHash).ToPostRecordResult()
    post = PostRecordResultToResponseContent(postRecordResult)
    post.RepliesLength = postRepliesRecordExecutor.GetPostRepliesRecordCount(postHash)
    postRewardsRecord := postRewardsRecordExecutor.GetPostRewardsRecordByPostHash(postHash)
    post.DeltaFuel = postRewardsRecord.DeltaFuel
    post.DeltaReputation = postRewardsRecord.DeltaReputation
    post.DeltaMilestonePoints = postRewardsRecord.DeltaMilestonePoints
    post.WithdrawableMPs= postRewardsRecord.WithdrawableMPs
    posts[index] = *post
    log.Printf("Post Content is loaded for postHash %s\n", postHash)
  }
  response.Posts = &posts
  log.Printf("Post Content for all postHashes are loaded\n")

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
  // Actor: "0x009",
  //}
  //response, _ := Handler(request)
  //fmt.Printf("%+v", response)

  lambda.Start(Handler)
}
