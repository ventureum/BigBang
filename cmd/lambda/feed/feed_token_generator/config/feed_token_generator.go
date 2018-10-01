package config

import (
  "gopkg.in/GetStream/stream-go2.v1"
  "os"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/error_config"
  "log"
)


type Request struct {
  FeedSlug string `json:"feedSlug,required"`
  UserId string `json:"userId,required"`
  GetStreamApiKey string `json:"getStreamApiKey,omitEmpty"`
  GetStreamApiSecret string `json:"getStreamApiSecret,omitEmpty"`
}

type Response struct {
  FeedToken string `json:"feedToken,omitempty"`
  Ok bool `json:"ok"`
  Message *error_config.ErrorInfo `json:"message,omitempty"`
}

func ProcessRequest(request Request, response *Response) {
  defer func() {
   if errPanic := recover(); errPanic != nil { //catch
     log.Printf("errPanic %+v", errPanic)
     response.FeedToken  = ""
     response.Message = error_config.CreatedErrorInfoFromString(errPanic)
   }
  }()

  var client *stream.Client
  var err error
  if request.GetStreamApiKey != "" && request.GetStreamApiSecret != "" {
    client, err = stream.NewClient(request.GetStreamApiKey, request.GetStreamApiSecret)
  } else {
    client, err = stream.NewClientFromEnv()
    request.GetStreamApiSecret = os.Getenv("STREAM_API_SECRET")
    request.GetStreamApiKey = os.Getenv("STREAM_API_KEY")
  }

  if err != nil {
   errorInfo := error_config.ErrorInfo{
     ErrorCode: error_config.GetStreamClientConnectionError,
   }
   log.Printf("Failed to connect GetStream client with error: %+v\n", err)
   log.Panicln(errorInfo.Marshal())
  }

  client.FlatFeed(request.FeedSlug, request.UserId)
  feedID := feed_attributes.CreateFeedId(request.FeedSlug, request.UserId)
  response.FeedToken = feedID.FeedToken(request.GetStreamApiSecret)
  response.Ok = true
}


func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
