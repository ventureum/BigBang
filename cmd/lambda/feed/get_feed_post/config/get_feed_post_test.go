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
        PostHash: "0xTestPostHash001",
        Requestor: "0xLambdaProfileActor001",
      },

      response: Response {
        Post: &ResponseContent {
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
            Meta: "[{offset: 0, length: 41, type: 'url'}]",
          },
          DeltaFuel: -50,
          DeltaReputation: 0,
          DeltaMilestonePoints: 0,
          WithdrawableMPs: 0,
          RepliesLength: 0,
        },
        PostVoteCountInfo: &feed_attributes.VoteCountInfo{
          DownVoteCount: 0,
          UpVoteCount: 0,
          TotalVoteCount: 0,
        },
        RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
          DownVoteCount: 0,
          UpVoteCount: 0,
          TotalVoteCount: 0,
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
