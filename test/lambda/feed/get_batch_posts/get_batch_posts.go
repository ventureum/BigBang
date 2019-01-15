package get_batch_posts_test

import (
	"BigBang/cmd/lambda/feed/get_batch_posts/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_batch_posts_config.Request
		response lambda_get_batch_posts_config.Response
		err      error
	}{
		{
			request: lambda_get_batch_posts_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_batch_posts_config.RequestContent{
					PostHashes: []string{
						test_constants.PostHash1,
						test_constants.PostHash2,
					},
				},
			},

			response: lambda_get_batch_posts_config.Response{
				Posts: &[]lambda_get_batch_posts_config.ResponseContent{
					{
						Actor:                test_constants.Actor1,
						Username:             test_constants.UserName1,
						PhotoUrl:             "http://123.com",
						BoardId:              test_constants.BoardId1,
						ParentHash:           test_constants.EmptyParentHash,
						PostHash:             test_constants.PostHash1,
						PostType:             string(feed_attributes.PostPostType),
						Content:              &test_constants.PostContent1,
						DeltaFuel:            -50,
						DeltaReputation:      0,
						DeltaMilestonePoints: 0,
						WithdrawableMPs:      0,
						RepliesLength:        0,
					},
					{
						Actor:                test_constants.Actor1,
						Username:             test_constants.UserName1,
						PhotoUrl:             "http://123.com",
						BoardId:              test_constants.BoardId1,
						ParentHash:           test_constants.EmptyParentHash,
						PostHash:             test_constants.PostHash2,
						PostType:             string(feed_attributes.PostPostType),
						Content:              &test_constants.PostContent2,
						DeltaFuel:            -50,
						DeltaReputation:      0,
						DeltaMilestonePoints: 0,
						WithdrawableMPs:      0,
						RepliesLength:        0,
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_batch_posts_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
