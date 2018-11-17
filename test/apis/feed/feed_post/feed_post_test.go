package feed_post

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/feed_post/config"
  "BigBang/test/constants"
  "BigBang/internal/pkg/api"
  "github.com/mitchellh/mapstructure"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_feed_post_config.Request
    response lambda_feed_post_config.Response
    err    error
  }{
    {
      request: lambda_feed_post_config.Request {
        Actor: test_constants.Actor1,
        BoardId: test_constants.BoardId1,
        ParentHash: test_constants.EmptyParentHash,
        PostHash: test_constants.PostHash1,
        TypeHash: test_constants.PostTypeHash,
        Content: test_constants.PostContent1,
      },
      response: lambda_feed_post_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_feed_post_config.Request {
        Actor: test_constants.Actor1,
        BoardId: test_constants.BoardId1,
        ParentHash: test_constants.EmptyParentHash,
        PostHash: test_constants.PostHash2,
        TypeHash: test_constants.PostTypeHash,
        Content: test_constants.PostContent2,
      },
      response: lambda_feed_post_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    responseMessageFeedPost := api.SendPost(test.request, api.FeedPostAlphaEndingPoint)
    var responseFeedPost lambda_feed_post_config.Response
    mapstructure.Decode(*responseMessageFeedPost, &responseFeedPost)
    assert.Equal(t, test.response, responseFeedPost)
  }
}
