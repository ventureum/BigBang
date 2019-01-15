package get_validator_recent_rating_vote_activities_test

import (
	"BigBang/cmd/lambda/TCR/get_validator_recent_rating_vote_activities/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/internal/platform/postgres_config/TCR/rating_vote_config"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	emptyRatingVoteActivities := []tcr_attributes.RatingVoteActivity{}

	tests := []struct {
		request  lambda_get_validator_recent_rating_vote_activities_config.Request
		response lambda_get_validator_recent_rating_vote_activities_config.Response
		err      error
	}{
		{
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor6,
					Limit: 0,
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &emptyRatingVoteActivities,
					NextCursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor6, test_constants.RatingVoteBlockTimestamp5),
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 5,
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &[]tcr_attributes.RatingVoteActivity{
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         20,
							Weight:         30,
						},
					},
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 4,
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &[]tcr_attributes.RatingVoteActivity{
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         20,
							Weight:         30,
						},
					},
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 3,
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &[]tcr_attributes.RatingVoteActivity{
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         30,
							Weight:         40,
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
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 3,
					Cursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId2, 1, 1, test_constants.Actor2, test_constants.RatingVoteBlockTimestamp3),
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &[]tcr_attributes.RatingVoteActivity{
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         20,
							Weight:         30,
						},
					},
					NextCursor: "",
				},
				Ok: true,
			},
			err: nil,
		},
		{
			request: lambda_get_validator_recent_rating_vote_activities_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_validator_recent_rating_vote_activities_config.RequestContent{
					Actor: test_constants.Actor2,
					Limit: 3,
					Cursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId2, 1, 2, test_constants.Actor2, test_constants.RatingVoteBlockTimestamp4),
				},
			},
			response: lambda_get_validator_recent_rating_vote_activities_config.Response{
				ResponseData: &lambda_get_validator_recent_rating_vote_activities_config.ResponseData{
					RatingVoteActivities: &[]tcr_attributes.RatingVoteActivity{
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp4,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId002",
							MilestoneId:    1,
							ObjectiveId:    1,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp3,
							Rating:         30,
							Weight:         40,
						},
						{
							ProjectId:      "0xProjectId001",
							MilestoneId:    1,
							ObjectiveId:    2,
							BlockTimestamp: test_constants.RatingVoteBlockTimestamp2,
							Rating:         30,
							Weight:         40,
						},
					},
					NextCursor: rating_vote_config.GenerateEncodedRatingVoteRecordID(
						test_constants.ProjectId1, 1, 1, test_constants.Actor2, test_constants.RatingVoteBlockTimestamp2),
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_validator_recent_rating_vote_activities_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.ResponseData.NextCursor, result.ResponseData.NextCursor)
		assert.Equal(t, test.response.ResponseData.RatingVoteActivities, result.ResponseData.RatingVoteActivities)
	}
}
