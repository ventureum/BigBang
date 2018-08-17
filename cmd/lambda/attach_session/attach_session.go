package main

import (

  "log"
  "gopkg.in/GetStream/stream-go2.v1"

  "github.com/aws/aws-lambda-go/lambda"

  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/session_record_config"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/getstream_config"
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
  Message string `json:"message,omitempty"`
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
  sessionRecord := request.ToSessionRecord()
  defer postgresFeedClient.Close()

  postgresFeedClient.Begin()
  sessionRecordExecutor := session_record_config.SessionRecordExecutor{*postgresFeedClient}
  updatedTimestamp := sessionRecordExecutor.UpsertSessionRecordTx(sessionRecord)

  // Insert Activity to GetStream
  activity := sessionRecord.ConvertSessionRecordToActivity(feed_attributes.BlockTimestamp(updatedTimestamp.Unix()))
  getStreamClient.AddFeedActivityToGetStream(activity)

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
  //content := feed_attributes.Content{
  //  Title: "titleSample1",
  //  Text: "hello, world",
  //}
  //request := Request{
  //  Actor:  "0x00999",
  //  PostHash: "0x009",
  //  StartTime: 1213444,
  //  EndTime:   1213778,
  //  Content: content,
  //}
  //reposnse, _ := Handler(request)
  //log.Printf("%+v", reposnse)

  lambda.Start(Handler)
}
