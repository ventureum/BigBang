package feed_token_generator_test

import (
	"BigBang/cmd/lambda/feed/feed_token_generator/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_feed_token_generator_config.Request
		response lambda_feed_token_generator_config.Response
		err      error
	}{
		{
			request: lambda_feed_token_generator_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_feed_token_generator_config.RequestContent{
					FeedSlug: "User",
					UserId:   "david3620",
				},
			},
			response: lambda_feed_token_generator_config.Response{
				FeedToken: "8XDj7VcxGoOYMYigN_bIT7h9hAo",
				Ok:        true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_feed_token_generator_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
