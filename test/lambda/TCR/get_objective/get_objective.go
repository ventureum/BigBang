package get_objective_test

import (
	"BigBang/cmd/lambda/TCR/get_objective/config"
	"BigBang/internal/app/tcr_attributes"
	"BigBang/test/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request  lambda_get_objective_config.Request
		response lambda_get_objective_config.Response
		err      error
	}{
		{
			request: lambda_get_objective_config.Request{
				PrincipalId: test_constants.Actor1,
				Body: lambda_get_objective_config.RequestContent{
					ProjectId:   test_constants.ProjectId1,
					MilestoneId: test_constants.MilestoneId1,
					ObjectiveId: test_constants.ObjectiveId1,
				},
			},
			response: lambda_get_objective_config.Response{
				Objective: &tcr_attributes.Objective{
					ProjectId:      test_constants.ProjectId1,
					MilestoneId:    test_constants.MilestoneId1,
					ObjectiveId:    test_constants.ObjectiveId1,
					Content:        test_constants.ObjectiveContent1,
					BlockTimestamp: test_constants.ObjectiveBlockTimestamp1,
				},
				Ok: true,
			},
			err: nil,
		},
	}
	for _, test := range tests {
		result, err := lambda_get_objective_config.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.response.Ok, result.Ok)
		assert.Equal(t, test.response.Objective.ProjectId, result.Objective.ProjectId)
		assert.Equal(t, test.response.Objective.MilestoneId, result.Objective.MilestoneId)
		assert.Equal(t, test.response.Objective.ObjectiveId, result.Objective.ObjectiveId)
		assert.Equal(t, test.response.Objective.Content, result.Objective.Content)
		assert.Equal(t, test.response.Objective.BlockTimestamp, result.Objective.BlockTimestamp)
		assert.Equal(t, test.response.Objective.AvgRating, result.Objective.AvgRating)
	}
}
