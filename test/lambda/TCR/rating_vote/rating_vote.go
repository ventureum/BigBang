package rating_vote_test

import (
	"BigBang/cmd/lambda/TCR/rating_vote/config"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/internal/platform/postgres_config/client_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_rating_vote_config.Request
		response lambda_rating_vote_config.Response
		err      error
	}{
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor1,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp1,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor2,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor3,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor4,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor5,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp5,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor6,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp5,
					Rating:         20,
					Weight:         30,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    2,
					Voter:          test_constants.Actor1,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
					Rating:         25,
					Weight:         35,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    1,
					ObjectiveId:    2,
					Voter:          test_constants.Actor2,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
					Rating:         30,
					Weight:         40,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    1,
					ObjectiveId:    1,
					Voter:          test_constants.Actor2,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
					Rating:         30,
					Weight:         40,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_rating_vote_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_rating_vote_config.RequestContent{
					ProjectId:      test_constants.ProjectId2,
					MilestoneId:    1,
					ObjectiveId:    2,
					Voter:          test_constants.Actor2,
					BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
					Rating:         30,
					Weight:         40,
				},
			},
			response: lambda_rating_vote_config.Response{
				Ok: true,
			},
			err: nil,
		},
	}
	postgresBigBangClient := client_config.ConnectPostgresClient(nil)
	ratingVoteExecutor := rating_vote_config.RatingVoteExecutor{*postgresBigBangClient}
	ratingVoteExecutor.ClearRatingVoteTable()

	for _, test := range tests {
		result, err := lambda_rating_vote_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
