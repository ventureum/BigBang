package feed_redeem_milestone_points_test

import (
	"BigBang/cmd/lambda/feed/feed_redeem_milestone_points/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_feed_redeem_milestone_points_config.Request
		response lambda_feed_redeem_milestone_points_config.Response
		err      error
	}{
		{
			request: lambda_feed_redeem_milestone_points_config.Request{},
			response: lambda_feed_redeem_milestone_points_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_feed_redeem_milestone_points_config.Request{
				IncreasedRedeemBlockNum: 1,
			},
			response: lambda_feed_redeem_milestone_points_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_feed_redeem_milestone_points_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
