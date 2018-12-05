package get_recent_posts

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
  "BigBang/cmd/lambda/feed/get_recent_posts/config"
  "BigBang/test/constants"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_get_recent_posts_config.Request
    response lambda_get_recent_posts_config.Response
    err    error
  }{
    {
      request: lambda_get_recent_posts_config.Request {
        Actor: test_constants.Actor1,
        TypeHash: test_constants.PostTypeHash,
      },
      response: lambda_get_recent_posts_config.Response {
        RecentPosts: &[]post_rewards_record_config.PostRewardsRecord{
          {
            PostHash: test_constants.PostHash2,
            Actor: test_constants.Actor1,
            PostType: string(feed_attributes.PostPostType),
            DeltaFuel: -50,
            DeltaReputation: 0,
            DeltaMilestonePoints: 0,
            WithdrawableMPs: 0,
          },
          {
            PostHash: test_constants.PostHash1,
            Actor: test_constants.Actor1,
            PostType: string(feed_attributes.PostPostType),
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
    responseMessageGetRecentPosts := api.SendPost(test.request, api.GetRecentPostsAlphaEndingPoint)
    var responseGetRecentPosts lambda_get_recent_posts_config.Response
    responseGetRecentPosts.RecentPosts = &[]post_rewards_record_config.PostRewardsRecord{}
    mapstructure.Decode(*responseMessageGetRecentPosts, &responseGetRecentPosts)
    mapstructure.Decode((*responseMessageGetRecentPosts).(map[string]interface{})["recentPosts"], responseGetRecentPosts.RecentPosts)
    assert.Equal(t, test.response.Ok, responseGetRecentPosts.Ok)
    resultRecentPosts := *(responseGetRecentPosts.RecentPosts)
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
