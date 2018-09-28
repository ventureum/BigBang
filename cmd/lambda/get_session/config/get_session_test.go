package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/session_record_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        PostHash: "0xTestPostHash002",
      },
      response: Response {
        Session: &session_record_config.SessionRecordResult{
          Actor: "0xLambdaProfileActor001",
          PostHash: "0xTestPostHash002",
          StartTime: 13000,
          EndTime: 14345,
          Content: &feed_attributes.Content{
            Title: "Title2",
            Text: "Text2",
            Image: "Image2",
            Subtitle: "Subtitle2",
          },
        },
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response.Session.Actor, result.Session.Actor)
    assert.Equal(t, test.response.Session.PostHash, result.Session.PostHash)
    assert.Equal(t, test.response.Session.StartTime, result.Session.StartTime)
    assert.Equal(t, test.response.Session.EndTime, result.Session.EndTime)
  }
}
