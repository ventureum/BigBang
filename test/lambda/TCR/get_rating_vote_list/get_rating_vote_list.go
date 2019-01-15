package get_rating_vote_list_test

import (
	"BigBang/cmd/lambda/TCR/get_rating_vote_list/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_rating_vote_list_config.Request
		response lambda_get_rating_vote_list_config.Response
		err      error
	}{
		{
			request: lambda_get_rating_vote_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_rating_vote_list_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: 1,
					ObjectiveId: 1,
					Limit:       0,
				},
			},
			response: lambda_get_rating_vote_list_config.Response{
				ResponseData: &lambda_get_rating_vote_list_config.ResponseData{
					ObjectiveVotesInfo: &tcr_attributes.ObjectiveVotesInfo{
						ProjectId:   test_constants.ProjectId1,
						MilestoneId: 1,
						ObjectiveId: 1,
						RatingVotes: &[]tcr_attributes.RatingVote{},
					},
					NextCursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor6, test_constants.RatingVoteBlockTimestamp5),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_rating_vote_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_rating_vote_list_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: 1,
					ObjectiveId: 1,
					Limit:       2,
				},
			},
			response: lambda_get_rating_vote_list_config.Response{
				ResponseData: &lambda_get_rating_vote_list_config.ResponseData{
					ObjectiveVotesInfo: &tcr_attributes.ObjectiveVotesInfo{
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
						},
					},
					NextCursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor4, test_constants.RatingVoteBlockTimestamp4),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_rating_vote_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_rating_vote_list_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: 1,
					ObjectiveId: 1,
					Limit:       2,
					Cursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor4, test_constants.RatingVoteBlockTimestamp4),
				},
			},
			response: lambda_get_rating_vote_list_config.Response{
				ResponseData: &lambda_get_rating_vote_list_config.ResponseData{
					ObjectiveVotesInfo: &tcr_attributes.ObjectiveVotesInfo{
						ProjectId:   test_constants.ProjectId1,
						MilestoneId: 1,
						ObjectiveId: 1,
						RatingVotes: &[]tcr_attributes.RatingVote{
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
						},
					},
					NextCursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor2, test_constants.RatingVoteBlockTimestamp2),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_rating_vote_list_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_rating_vote_list_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: 1,
					ObjectiveId: 1,
					Limit:       5,
					Cursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor4, test_constants.RatingVoteBlockTimestamp4),
				},
			},
			response: lambda_get_rating_vote_list_config.Response{
				ResponseData: &lambda_get_rating_vote_list_config.ResponseData{
					ObjectiveVotesInfo: &tcr_attributes.ObjectiveVotesInfo{
						ProjectId:   test_constants.ProjectId1,
						MilestoneId: 1,
						ObjectiveId: 1,
						RatingVotes: &[]tcr_attributes.RatingVote{
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
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_rating_vote_list_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.ResponseData.NextCursor, result.ResponseData.NextCursor)
		objVoteInfo := *result.ResponseData.ObjectiveVotesInfo
		responseObjVoteInfo := *test.response.ResponseData.ObjectiveVotesInfo
		assert.Equal(t, responseObjVoteInfo.ProjectId, objVoteInfo.ProjectId)
		assert.Equal(t, responseObjVoteInfo.MilestoneId, objVoteInfo.MilestoneId)
		assert.Equal(t, responseObjVoteInfo.ObjectiveId, objVoteInfo.ObjectiveId)
		responseRatingVotes := *responseObjVoteInfo.RatingVotes
		for index, ratingVote := range *objVoteInfo.RatingVotes {
			assert.Equal(t, responseRatingVotes[index].Voter, ratingVote.Voter)
			assert.Equal(t, responseRatingVotes[index].Rating, ratingVote.Rating)
			assert.Equal(t, responseRatingVotes[index].Weight, ratingVote.Weight)
		}
	}
}
