package main

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
        PostHashes: []string{
          "0xTestPostHash001",
          "0xTestPostHash002",
        },
      },

      response: Response {
        Posts: &[]ResponseContent{
          {
            Actor: "0xLambdaProfileActor001",
            Username: "LambdaProfileActor001",
            PhotoUrl: "http://123.com",
            BoardId: "0xTestBoard001",
            ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
            PostHash: "0xTestPostHash001",
            PostType: "POST",
            Content: &feed_attributes.Content{
              Title: "Title1",
              Text: "Text1",
              Image: "Image1",
              Subtitle: "Subtitle1",
            },
            DeltaFuel: -50,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            WithdrawableMPs: 0,
            RepliesLength: 0,

          },
          {
            Actor: "0xLambdaProfileActor001",
            Username: "LambdaProfileActor001",
            PhotoUrl: "http://123.com",
            BoardId: "0xTestBoard001",
            ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
            PostHash: "0xTestPostHash002",
            PostType: "POST",
            Content: &feed_attributes.Content{
              Title: "Title2",
              Text: "Text2",
              Image: "Image2",
              Subtitle: "Subtitle2",
            },
            DeltaFuel: -50,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            WithdrawableMPs: 0,
            RepliesLength: 0,
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
    assert.Equal(t, test.response, result)
  }
}
