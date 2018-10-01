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
        BoardId: "0xTestBoard001",
        ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        PostHash: "0xTestPostHash001",
        TypeHash: "0x2fca5a5e",
        Content: feed_attributes.Content{
          Title: "Title1",
          Text: "Text1",
          Image: "Image1",
          Subtitle: "Subtitle1",
          Meta: "[{offset: 0, length: 41, type: 'url'}]",
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
        BoardId: "0xTestBoard001",
        ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
        PostHash: "0xTestPostHash002",
        TypeHash: "0x2fca5a5e",
        Content: feed_attributes.Content{
          Title: "Title2",
          Text: "Text2",
          Image: "Image2",
          Subtitle: "Subtitle2",
          Meta: "[{offset: 0, length: 41, type: 'url'}]",
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
