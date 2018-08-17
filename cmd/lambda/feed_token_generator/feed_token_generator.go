package main

import (
  "gopkg.in/GetStream/stream-go2.v1"
  "os"
  "github.com/aws/aws-lambda-go/lambda"
  "BigBang/internal/app/feed_attributes"
)


type Request struct {
  FeedSlug string `json:"feedSlug,required"`
  UserId string `json:"userId,required"`
  GetStreamApiKey string `json:"getStreamApiKey,omitEmpty"`
  GetStreamApiSecret string `json:"getStreamApiSecret,omitEmpty"`
}

type Response struct {
  FeedToken string `json:"feedToken"`
  Ok bool `json:"ok"`
}


func Handler(request Request) (Response, error) {
  var client *stream.Client
  var err error
  if request.GetStreamApiKey != "" && request.GetStreamApiSecret != "" {
    client, err = stream.NewClient(request.GetStreamApiKey, request.GetStreamApiSecret)
  } else {
    client, err = stream.NewClientFromEnv()
    request.GetStreamApiSecret = os.Getenv("STREAM_API_SECRET")
    request.GetStreamApiKey = os.Getenv("STREAM_API_KEY")
  }

  response := Response {
    FeedToken: "",
    Ok: false,
  }
  if err != nil {
    return response, err
  }
  client.FlatFeed(request.FeedSlug, request.UserId)
  feedID := feed_attributes.CreateFeedId(request.FeedSlug, request.UserId)
  response.FeedToken = feedID.FeedToken(request.GetStreamApiSecret)
  response.Ok = true
  return response, nil
}

func main() {
  lambda.Start(Handler)
}
