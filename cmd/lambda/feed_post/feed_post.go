package main

import (
  "log"
  "gopkg.in/GetStream/stream-go2.v1"
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/platform/eth_config"
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
  Message string `json:"message,omitempty"`
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
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      response.Message = errStr.(string)
    }
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

  postgresFeedClient := client_config.ConnectPostgresClient()
  getStreamClient := &getstream_config.GetStreamClient{C: getStreamIOClient}
  defer postgresFeedClient.Close()
  postRecord := request.ToPostRecord()
  eth_config.ProcessPostRecord(
    postRecord,
    getStreamClient,
    postgresFeedClient,
    feed_attributes.OFF_CHAIN)
  response.Ok = true
}

func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}

func main() {
  // TODO(david.shao): remove example when deployed to production
  //content := feed_attributes.Content{
  //  Title: "titleSample1",
  //  Text: "hello, world",
  //}
  //request := Request{
  //  Actor:  "0x00999",
  //  BoardId: "0x02",
  //  ParentHash: "0x007",
  //  PostHash: "0x009",
  //  TypeHash:  feed_attributes.ReplyPostType.Hash(),
  //  Content: content,
  //}
  //reposnse, _ := Handler(request)
  //log.Printf("%+v", reposnse)

  lambda.Start(Handler)
}
