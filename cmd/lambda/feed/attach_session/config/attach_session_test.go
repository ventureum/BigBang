package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/app/feed_attributes"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor001",
        PostHash: "0xTestPostHash001",
        StartTime: 10000,
        EndTime: 12345,
        Content: feed_attributes.Content{
          Title: "Title1",
          Text: "Text1",
          Image: "Image1",
          Subtitle: "Subtitle1",
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: Request {
        Actor: "0xLambdaProfileActor001",
        PostHash: "0xTestPostHash002",
        StartTime: 13000,
        EndTime: 14345,
        Content: feed_attributes.Content{
          Title: "Title2",
          Text: "Text2",
          Image: "Image2",
          Subtitle: "Subtitle2",
        },
      },
      response: Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
