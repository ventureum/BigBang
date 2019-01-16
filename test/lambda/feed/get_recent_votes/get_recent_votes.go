package get_recent_votes_test

import (
	"BigBang/cmd/lambda/feed/get_recent_votes/config"
	"BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_recent_votes_config.Request
		response lambda_get_recent_votes_config.Response
		err      error
	}{
		{
			request: lambda_get_recent_votes_config.Request{
				PrincipalId: test_constants.Actor2,
				Body: lambda_get_recent_votes_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 20,
				},
			},
			response: lambda_get_recent_votes_config.Response{
				RecentVotes: &[]post_votes_record_config.PostVotesRecord{
					{
						PostHash:             test_constants.PostHash1,
						Actor:                test_constants.Actor2,
						PostType:             "POST",
						VoteType:             "UP",
						DeltaFuel:            -18,
						DeltaReputation:      0,
						DeltaMilestonePoints: 0,
						SignedReputation:     100,
					},
					{
						PostHash:             test_constants.PostHash1,
						Actor:                test_constants.Actor2,
						PostType:             "POST",
						VoteType:             "DOWN",
						DeltaFuel:            -20,
						DeltaReputation:      0,
						DeltaMilestonePoints: 0,
						SignedReputation:     -100,
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_recent_votes_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		resultRecentPosts := *result.RecentVotes
		responseRecentPosts := *test.response.RecentVotes
		assert.Equal(t, len(resultRecentPosts), len(responseRecentPosts))
		for index, responseRecentPost := range responseRecentPosts {
			assert.Equal(t, resultRecentPosts[index].Actor, responseRecentPost.Actor)
			assert.Equal(t, resultRecentPosts[index].PostHash, responseRecentPost.PostHash)
			assert.Equal(t, resultRecentPosts[index].PostType, responseRecentPost.PostType)
			assert.Equal(t, resultRecentPosts[index].VoteType, responseRecentPost.VoteType)
			assert.Equal(t, resultRecentPosts[index].DeltaFuel, responseRecentPost.DeltaFuel)
			assert.Equal(t, resultRecentPosts[index].DeltaReputation, responseRecentPost.DeltaReputation)
			assert.Equal(t, resultRecentPosts[index].DeltaMilestonePoints, responseRecentPost.DeltaMilestonePoints)
			assert.Equal(t, resultRecentPosts[index].SignedReputation, responseRecentPost.SignedReputation)
		}
	}
}
