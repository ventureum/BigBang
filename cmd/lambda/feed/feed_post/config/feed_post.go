package lambda_feed_post_config

import (
  "log"
  "gopkg.in/GetStream/stream-go2.v1"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/platform/eth_config"
  "BigBang/internal/pkg/error_config"
)


type Request struct {
  Actor string `json:"actor,required"`
  BoardId string `json:"boardId,required"`
  ParentHash string `json:"parentHash,required"`
  PostHash string `json:"postHash,required"`
  TypeHash string `json:"typeHash,required"`
  Content feed_attributes.Content `json:"content,required"`
  GetStreamApiKey string `json:"getStreamApiKey,omitempty"`
  GetStreamApiSecret string `json:"getStreamApiSecret,omitempty"`
}

type Response struct {
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func (request *Request) ToPostRecord() (*post_config.PostRecord) {
  return &post_config.PostRecord{
    Actor:      request.Actor,
    BoardId:    request.BoardId,
    ParentHash: request.ParentHash,
    PostHash:   request.PostHash,
    PostType:   feed_attributes.CreatePostTypeFromHashStr(request.TypeHash).Value(),
    Content:    request.Content.ToJsonText(),
  }
}

func ProcessRequest(request Request, response *Response) {
  postgresBigBangClient := client_config.ConnectPostgresClient()
  defer func() {
    if errPanic := recover(); errPanic != nil { //catch
      response.Message = error_config.CreatedErrorInfoFromString(errPanic)
      postgresBigBangClient.RollBack()
    }
    postgresBigBangClient.Close()
  }()

  var err error
  var getStreamIOClient *stream.Client
  if request.GetStreamApiKey != "" && request.GetStreamApiSecret != "" {
    getStreamIOClient, err = stream.NewClient(request.GetStreamApiKey, request.GetStreamApiSecret)
  } else {
    getStreamIOClient, err = stream.NewClientFromEnv()
  }

  if err != nil {
    log.Panic(err.Error())
  }

  getStreamClient := &getstream_config.GetStreamClient{C: getStreamIOClient}

  postRecord := request.ToPostRecord()
  eth_config.ProcessPostRecord(
    postRecord,
    getStreamClient,
    postgresBigBangClient,
    feed_attributes.OFF_CHAIN)
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
