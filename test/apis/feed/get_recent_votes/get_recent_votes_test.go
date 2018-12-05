package config

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/cmd/lambda/feed/get_recent_votes/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_recent_votes_config.Request
    response lambda_get_recent_votes_config.Response
    err    error
  }{
    {
      request: lambda_get_recent_votes_config.Request {
        Actor: test_constants.Actor2,
        Limit: 20,
      },
      response: lambda_get_recent_votes_config.Response {
        RecentVotes: &[]post_votes_record_config.PostVotesRecord{
          {
            PostHash: test_constants.PostHash1,
            Actor: test_constants.Actor2,
            PostType: "POST",
            VoteType: "UP",
            DeltaFuel: -18,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            SignedReputation: 100,
          },
          {
            PostHash: test_constants.PostHash1,
            Actor: test_constants.Actor2,
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
    responseMessageGetRecentVotes := api.SendPost(test.request, api.GetRecentVotesAlphaEndingPoint)
    var responseGetRecentVotes lambda_get_recent_votes_config.Response
    mapstructure.Decode(*responseMessageGetRecentVotes , &responseGetRecentVotes)
    assert.Equal(t, test.response, responseGetRecentVotes)
    assert.Equal(t, test.response.Ok, responseGetRecentVotes.Ok)
    resultRecentVotes := *(responseGetRecentVotes.RecentVotes)
    responseRecentVotes := *(test.response.RecentVotes)
    assert.Equal(t, len(resultRecentVotes), len(responseRecentVotes))
    for index, responseRecentVote := range responseRecentVotes {
      assert.Equal(t, resultRecentVotes[index].Actor, responseRecentVote.Actor)
      assert.Equal(t, resultRecentVotes[index].PostHash, responseRecentVote.PostHash)
      assert.Equal(t, resultRecentVotes[index].PostType, responseRecentVote.PostType)
      assert.Equal(t, resultRecentVotes[index].VoteType, responseRecentVote.VoteType)
      assert.Equal(t, resultRecentVotes[index].DeltaFuel, responseRecentVote.DeltaFuel)
      assert.Equal(t, resultRecentVotes[index].DeltaReputation, responseRecentVote.DeltaReputation)
      assert.Equal(t, resultRecentVotes[index].DeltaMilestonePoints, responseRecentVote.DeltaMilestonePoints)
      assert.Equal(t, resultRecentVotes[index].SignedReputation, responseRecentVote.SignedReputation)
    }
  }
}
