package lambda_feed_token_generator_config

import (
  "gopkg.in/GetStream/stream-go2.v1"
  "os"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/error_config"
  "log"
  "BigBang/cmd/lambda/common/auth"
)

type Request struct {
  PrincipalId string `json:"principalId,required"`
  Body RequestContent `json:"body,required"`
}

type RequestContent struct {
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
     response.FeedToken  = ""
     response.Message = error_config.CreatedErrorInfoFromString(errPanic)
   }
  }()

  auth.AuthProcess(request.PrincipalId, "", nil)

  var client *stream.Client
  var err error
  if request.Body.GetStreamApiKey != "" && request.Body.GetStreamApiSecret != "" {
    client, err = stream.NewClient(request.Body.GetStreamApiKey, request.Body.GetStreamApiSecret)
  } else {
    client, err = stream.NewClientFromEnv()
    request.Body.GetStreamApiSecret = os.Getenv("STREAM_API_SECRET")
    request.Body.GetStreamApiKey = os.Getenv("STREAM_API_KEY")
  }

  if err != nil {
   errorInfo := error_config.ErrorInfo{
     ErrorCode: error_config.GetStreamClientConnectionError,
   }
   log.Printf("Failed to connect GetStream client with error: %+v\n", err)
   log.Panicln(errorInfo.Marshal())
  }

  client.FlatFeed(request.Body.FeedSlug, request.Body.UserId)
  feedID := feed_attributes.CreateFeedId(request.Body.FeedSlug, request.Body.UserId)
  response.FeedToken = feedID.FeedToken(request.Body.GetStreamApiSecret)
  response.Ok = true
}


func Handler(request Request) (response Response, err error) {
  response.Ok = false
  ProcessRequest(request, &response)
  return response, nil
}
