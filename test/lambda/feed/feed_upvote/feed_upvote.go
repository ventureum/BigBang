package feed_upvote_test

import (
	"BigBang/cmd/lambda/feed/feed_upvote/config"
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/error_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_feed_upvote_config.Request
		response lambda_feed_upvote_config.Response
		err      error
	}{
		{
			request: lambda_feed_upvote_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_feed_upvote_config.RequestContent{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					Value:    -1,
				},
			},
			response: lambda_feed_upvote_config.Response{
				VoteInfo: &feed_attributes.VoteInfo{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					FuelCost: 20,
					RewardsInfo: &feed_attributes.RewardsInfo{
						Fuel:            80,
						Reputation:      100,
						MilestonePoints: 100,
					},
					PostVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    0,
						TotalVoteCount: 1,
					},
					RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    0,
						TotalVoteCount: 1,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_feed_upvote_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_feed_upvote_config.RequestContent{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					Value:    0,
				},
			},
			response: lambda_feed_upvote_config.Response{
				VoteInfo: &feed_attributes.VoteInfo{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					FuelCost: 18,
					RewardsInfo: &feed_attributes.RewardsInfo{
						Fuel:            80,
						Reputation:      100,
						MilestonePoints: 100,
					},
					PostVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    0,
						TotalVoteCount: 1,
					},
					RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    0,
						TotalVoteCount: 1,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_feed_upvote_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_feed_upvote_config.RequestContent{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					Value:    1,
				},
			},
			response: lambda_feed_upvote_config.Response{
				VoteInfo: &feed_attributes.VoteInfo{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					FuelCost: 18,
					RewardsInfo: &feed_attributes.RewardsInfo{
						Fuel:            62,
						Reputation:      100,
						MilestonePoints: 100,
					},
					PostVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    1,
						TotalVoteCount: 2,
					},
					RequestorVoteCountInfo: &feed_attributes.VoteCountInfo{
						DownVoteCount:  1,
						UpVoteCount:    1,
						TotalVoteCount: 2,
					},
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_feed_upvote_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_feed_upvote_config.RequestContent{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					Value:    1,
				},
			},
			response: lambda_feed_upvote_config.Response{
				Message: &error_config.ErrorInfo{
					ErrorCode: "ExceedingUpvoteLimit",
					ErrorData: error_config.ErrorData{
						"actor":    test_constants.Actor2,
						"postHash": test_constants.PostHash1,
					},
					ErrorLocation: "ActorVotesCountersRecordLocation",
				},
				Ok: false,
			},
			err: nil,
		},
		{
			request: lambda_feed_upvote_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_feed_upvote_config.RequestContent{
					Actor:    test_constants.Actor2,
					PostHash: test_constants.PostHash1,
					Value:    -1,
				},
			},
			response: lambda_feed_upvote_config.Response{
				Message: &error_config.ErrorInfo{
					ErrorCode: "ExceedingDownvoteLimit",
					ErrorData: error_config.ErrorData{
						"actor":    test_constants.Actor2,
						"postHash": test_constants.PostHash1,
					},
					ErrorLocation: "ActorVotesCountersRecordLocation",
				},
				Ok: false,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_feed_upvote_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
