package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/post_votes_record_config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request Request
    response Response
    err    error
  }{
    {
      request: Request {
        Actor: "0xLambdaProfileActor002",
        Limit: 20,
      },
      response: Response {
        RecentVotes: &[]post_votes_record_config.PostVotesRecord{
          {
            PostHash: "0xTestPostHash001",
            Actor: "0xLambdaProfileActor002",
            PostType: "POST",
            VoteType: "UP",
            DeltaFuel: -10,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            SignedReputation: 100,
          },
          {
            PostHash: "0xTestPostHash001",
            Actor: "0xLambdaProfileActor002",
            PostType: "POST",
            VoteType: "DOWN",
            DeltaFuel: -20,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            SignedReputation: -100,
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
    resultRecentPosts := *result.RecentVotes
    responseRecentPosts := *test.response.RecentVotes
    assert.Equal(t, len(resultRecentPosts), len(responseRecentPosts))
    for index, responseRecentPost := range responseRecentPosts {
      assert.Equal(t, resultRecentPosts[index].Actor, responseRecentPost.Actor)
      assert.Equal(t, resultRecentPosts[index].PostHash, responseRecentPost.PostHash)
      assert.Equal(t, resultRecentPosts[index].PostType, responseRecentPost.PostType)
      assert.Equal(t, resultRecentPosts[index].VoteType, responseRecentPost.VoteType)
      assert.Equal(t, resultRecentPosts[index].DeltaFuel, responseRecentPost.DeltaFuel)
      assert.Equal(t, resultRecentPosts[index].DeltaReputation, responseRecentPost.DeltaReputation)
      assert.Equal(t, resultRecentPosts[index].DeltaMilestonePoints, responseRecentPost.DeltaMilestonePoints)
      assert.Equal(t, resultRecentPosts[index].SignedReputation, responseRecentPost.SignedReputation)
    }
  }
}
