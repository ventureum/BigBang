package lambda_get_batch_posts_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
)


type Request struct {
  PostHashes []string `json:"postHashes,required"`
}

type ResponseContent struct {
  Actor string `json:"actor,required"`
  Username string `json:"username,required"`
  PhotoUrl string `json:"photoUrl,required"`
  BoardId string `json:"boardId,required"`
  ParentHash string `json:"parentHash,required"`
  PostHash string `json:"postHash,required"`
  PostType string `json:"postType,required"`
  Content *feed_attributes.Content `json:"content,required"`
  DeltaFuel int64 `json:"deltaFuel,required"`
  DeltaReputation int64 `json:"deltaReputation,required"`
  DeltaMilestonePoints int64 `json:"deltaMilestonePoints,required"`
  WithdrawableMPs int64 `json:"withdrawableMPs,required"`
  RepliesLength int64 `json:"repliesLength,required"`
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
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Posts = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
    }
    postgresBigBangClient.Close()
  }()

  postHashes := request.PostHashes


  postExecutor := post_config.PostExecutor{*postgresBigBangClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}

  for _, postHash := range postHashes {
    postExecutor.VerifyPostRecordExisting(postHash)
  }
  posts := make([]ResponseContent, len(postHashes))
  for index, postHash := range postHashes {
    var post *ResponseContent
    postRecordResult := postExecutor.GetPostRecord(postHash).ToPostRecordResult()
    post = PostRecordResultToResponseContent(postRecordResult)
    actorProfileRecord := actorProfileRecordExecutor.GetActorProfileRecord(postRecordResult.Actor)
    post.Username = actorProfileRecord.Username
    post.PhotoUrl = actorProfileRecord.PhotoUrl
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
