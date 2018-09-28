package config

import (

  "log"
  "gopkg.in/GetStream/stream-go2.v1"

  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/pkg/error_config"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/eth_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  PostHash string `json:"postHash,required"`
  StartTime int64 `json:"startTime,required"`
  EndTime int64 `json:"endTime,required"`
  Content feed_attributes.Content `json:"content,required"`
  GetStreamApiKey string `json:"getStreamApiKey,omitempty"`
  GetStreamApiSecret string `json:"getStreamApiSecret,omitempty"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToSessionRecord() (*session_record_config.SessionRecord) {
  return &session_record_config.SessionRecord{
    Actor:      request.Actor,
    PostHash:   request.PostHash,
    StartTime:  request.StartTime,
    EndTime:    request.EndTime,
    Content:    request.Content.ToJsonText(),
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresFeedClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresFeedClient.RollBack()
    }
    postgresFeedClient.Close()
  }()

  var err error
  var getStreamIOClient *stream.Client
  if request.GetStreamApiKey != "" && request.GetStreamApiSecret != "" {
    getStreamIOClient, err = stream.NewClient(request.GetStreamApiKey, request.GetStreamApiSecret)
  } else {
    getStreamIOClient, err = stream.NewClientFromEnv()
  }

  if err != nil {
    log.Panic(err)
  }

  getStreamClient := &getstream_config.GetStreamClient{C: getStreamIOClient}
  sessionRecord := request.ToSessionRecord()

  postgresFeedClient.Begin()

  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}

  postExecutor.VerifyPostRecordExistingTx(sessionRecord.PostHash)

  postRecord := postExecutor.GetPostRecordTx(sessionRecord.PostHash)
  sessionRecordExecutor.UpsertSessionRecordTx(sessionRecord)

  activity := eth_config.ConvertPostRecordToActivity(
    postRecord, feed_attributes.OFF_CHAIN, feed_attributes.BlockTimestamp(postRecord.CreatedAt.Unix()))

  // Insert Activity to GetStream
  sessionRecord.EmbedSessionRecordToActivity(activity)
  getStreamClient.UpdateFeedActivityToGetStream(activity)

  postgresFeedClient.Commit()

  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
