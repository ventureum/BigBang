package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/post_rewards_record_config"
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
        TypeHash: "0x2fca5a5e",
      },
      response: Response {
        RecentPosts: &[]post_rewards_record_config.PostRewardsRecord{
          {
            PostHash: "0xTestPostHash002",
            Actor: "0xLambdaProfileActor001",
            PostType: "POST",
            DeltaFuel: -50,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            WithdrawableMPs: 0,
          },
          {
            PostHash: "0xTestPostHash001",
            Actor: "0xLambdaProfileActor001",
            PostType: "POST",
            DeltaFuel: -50,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            WithdrawableMPs: 0,
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
    assert.Equal(t, test.response.Ok, result.Ok)
    resultRecentPosts := *result.RecentPosts
    responseRecentPosts := *test.response.RecentPosts
    assert.Equal(t, len(resultRecentPosts), len(responseRecentPosts))
    for index, responseRecentPost := range responseRecentPosts {
      assert.Equal(t, resultRecentPosts[index].Actor, responseRecentPost.Actor)
      assert.Equal(t, resultRecentPosts[index].PostHash, responseRecentPost.PostHash)
      assert.Equal(t, resultRecentPosts[index].PostType, responseRecentPost.PostType)
      assert.Equal(t, resultRecentPosts[index].DeltaFuel, responseRecentPost.DeltaFuel)
      assert.Equal(t, resultRecentPosts[index].DeltaReputation, responseRecentPost.DeltaReputation)
      assert.Equal(t, resultRecentPosts[index].DeltaMilestonePoints, responseRecentPost.DeltaMilestonePoints)
      assert.Equal(t, resultRecentPosts[index].WithdrawableMPs, responseRecentPost.WithdrawableMPs)
    }
  }
}
