package get_milestone_test

import (
	"BigBang/cmd/lambda/TCR/get_milestone/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_milestone_config.Request
		response lambda_get_milestone_config.Response
		err      error
	}{
		{
			request: lambda_get_milestone_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_milestone_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: test_constants.MilestoneId1,
				},
			},
			response: lambda_get_milestone_config.Response{
				Milestone: &tcr_attributes.Milestone{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					Content:        test_constants.MilestoneContent1,
					BlockTimestamp: test_constants.MilestoneBlockTimestamp1,
					StartTime:      test_constants.MilestoneStartTime1,
					EndTime:        test_constants.MilestoneEndTime1,
					State:          tcr_attributes.CompleteMilestoneState,
					NumObjectives:  2,
					Objectives: &[]tcr_attributes.Objective{
						{
							ProjectId:      test_constants.ProjectId1,
							MilestoneId:    test_constants.MilestoneId1,
							ObjectiveId:    test_constants.ObjectiveId1,
							Content:        test_constants.ObjectiveContent1,
							BlockTimestamp: test_constants.ObjectiveBlockTimestamp1,
						},
						{
							ProjectId:      test_constants.ProjectId1,
							MilestoneId:    test_constants.MilestoneId1,
							ObjectiveId:    test_constants.ObjectiveId2,
							Content:        test_constants.ObjectiveContent1,
							BlockTimestamp: test_constants.ObjectiveBlockTimestamp2,
						},
					},
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_milestone_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.Milestone.ProjectId, result.Milestone.ProjectId)
		assert.Equal(t, test.response.Milestone.MilestoneId, result.Milestone.MilestoneId)
		assert.Equal(t, test.response.Milestone.Content, result.Milestone.Content)
		assert.Equal(t, test.response.Milestone.BlockTimestamp, result.Milestone.BlockTimestamp)
		assert.Equal(t, test.response.Milestone.StartTime, result.Milestone.StartTime)
		assert.Equal(t, test.response.Milestone.EndTime, result.Milestone.EndTime)
		assert.Equal(t, string(test.response.Milestone.State), string(result.Milestone.State))
		assert.Equal(t, test.response.Milestone.NumObjectives, result.Milestone.NumObjectives)
		assert.Equal(t, test.response.Milestone.NumObjectives, int64(len(*result.Milestone.Objectives)))
		assert.Equal(t, test.response.Milestone.Objectives, result.Milestone.Objectives)
	}
}
