package lambda_feed_upvote_config

import (
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/eth_config"
  "BigBang/internal/pkg/error_config"
)

type Request struct {
  Actor string `json:"actor,required"`
  PostHash string `json:"postHash,required"`
  Value int64 `json:"value,required"`
}

type Response struct {
  VoteInfo *feed_attributes.VoteInfo `json:"voteInfo,omitempty"`
  Ok      bool   `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient(nil)
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.VoteInfo = nil
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      if feed_attributes.CreateVoteTypeFromValue(request.Value) != feed_attributes.LOOKUP_VOTE_TYPE {
        postgresBigBangClient.RollBack()
      }
    }
    postgresBigBangClient.Close()
  }()

  postVotesRecord := post_votes_record_config.PostVotesRecord {
    Actor: request.Actor,
    PostHash: request.PostHash,
    VoteType: feed_attributes.CreateVoteTypeFromValue(request.Value),
  }

  if postVotesRecord.VoteType == feed_attributes.LOOKUP_VOTE_TYPE {
    response.VoteInfo =  eth_config.QueryPostVotesInfo(&postVotesRecord, postgresBigBangClient)
  } else {
    response.VoteInfo = eth_config.ProcessPostVotesRecord(&postVotesRecord, postgresBigBangClient)
  }

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
