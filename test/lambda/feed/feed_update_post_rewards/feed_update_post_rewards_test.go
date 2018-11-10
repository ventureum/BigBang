package feed_update_post_rewards

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "BigBang/cmd/lambda/feed/feed_update_post_rewards/config"
)

func TestHandler(t *testing.T) {
  tests := []struct{
    request lambda_feed_update_post_rewards_config.Request
    response lambda_feed_update_post_rewards_config.Response
    err    error
  }{
    {
      request: lambda_feed_update_post_rewards_config.Request {},
      response: lambda_feed_update_post_rewards_config.Response {
        Ok: true,
      },
      err: nil,
    },
    {
      request: lambda_feed_update_post_rewards_config.Request {
        IncreasedRedeemBlockNum: 1,
      },
      response: lambda_feed_update_post_rewards_config.Response {
        Ok: true,
      },
      err: nil,
    },
  }
  for _, test := range tests {
    result, err := lambda_feed_update_post_rewards_config.Handler(test.request)
    assert.IsType(t, test.err, err)
    assert.Equal(t, test.response, result)
  }
}
