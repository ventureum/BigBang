package feed_post_test

import (
	"BigBang/cmd/lambda/feed/feed_post/config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_feed_post_config.Request
		response lambda_feed_post_config.Response
		err      error
	}{
		{
			request: lambda_feed_post_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_feed_post_config.RequestContent{
					Actor:      test_constants.Actor1,
					BoardId:    test_constants.BoardId1,
					ParentHash: test_constants.EmptyParentHash,
					PostHash:   test_constants.PostHash1,
					TypeHash:   test_constants.PostTypeHash,
					Content:    test_constants.PostContent1,
				},
			},
			response: lambda_feed_post_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_feed_post_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_feed_post_config.RequestContent{
					Actor:      test_constants.Actor1,
					BoardId:    test_constants.BoardId1,
					ParentHash: test_constants.EmptyParentHash,
					PostHash:   test_constants.PostHash2,
					TypeHash:   test_constants.PostTypeHash,
					Content:    test_constants.PostContent2,
				},
			},
			response: lambda_feed_post_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_feed_post_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
