package get_batch_rating_vote_list_test

import (
	"BigBang/cmd/lambda/TCR/get_batch_rating_vote_list/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_batch_rating_vote_list_config.Request
		response lambda_get_batch_rating_vote_list_config.Response
		err      error
	}{
		{
			request: lambda_get_batch_rating_vote_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_batch_rating_vote_list_config.RequestContent{
					ObjectiveVotesInfoKeyList: []tcr_attributes.ObjectiveVotesInfoKey{
						{
							ProjectId:   test_constants.ProjectId1,
							MilestoneId: 1,
							ObjectiveId: 1,
						},
					},
				},
			},
			response: lambda_get_batch_rating_vote_list_config.Response{
				ObjectiveVotesInfoList: &[]tcr_attributes.ObjectiveVotesInfo{
					{
						ProjectId:   test_constants.ProjectId1,
						MilestoneId: 1,
						ObjectiveId: 1,
						RatingVotes: &[]tcr_attributes.RatingVote{
							{
								Voter:          test_constants.Actor6,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp5,
							},
							{
								Voter:          test_constants.Actor5,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp5,
							},
							{
								Voter:          test_constants.Actor4,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
							},
							{
								Voter:          test_constants.Actor3,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							},
							{
								Voter:          test_constants.Actor2,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							},
							{
								Voter:          test_constants.Actor1,
								Rating:         20,
								Weight:         30,
								BlockTimestamp: test_constants.RatingVoteBlockTimestamp1,
							},
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_batch_rating_vote_list_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response, result)
	}
}
