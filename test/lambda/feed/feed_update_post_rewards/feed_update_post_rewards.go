package feed_update_post_rewards_test

import (
	"BigBang/cmd/lambda/feed/feed_update_post_rewards/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		response lambda_feed_update_post_rewards_config.Response
		err      error
	}{
		{
			response: lambda_feed_update_post_rewards_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			response: lambda_feed_update_post_rewards_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_feed_update_post_rewards_config.Handler()
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
